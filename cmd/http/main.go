package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/devshark/alchemy-fleet/adapters/database"
	"github.com/devshark/alchemy-fleet/app"
	"github.com/devshark/alchemy-fleet/ent"
	_ "github.com/go-sql-driver/mysql"
)

const defaultTimeout = 5 * time.Minute
const shutdownTimeout = 5 * time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg := getConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)

	drv, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	defer drv.Close()

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("failed creating schema resources: ", err)
	}

	spacecraftService := database.New(client)

	server := app.NewHTTPServer(spacecraftService, cfg.Port, defaultTimeout)

	stop := make(chan os.Signal, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("ListenAndServe: ", err)
		}

		close(stop)
	}()

	signal.Notify(stop,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	log.Printf("Listening on port %d", cfg.Port)
	<-stop

	log.Print("Shutting down...")
	// if Shutdown takes longer than shutdownTimeout, cancel the context
	time.AfterFunc(shutdownTimeout, cancel)

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Shutdown on error: ", err)
	}

	log.Print("Server gracefully stopped")
}

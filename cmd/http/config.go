package main

import "github.com/spf13/viper"

type config struct {
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	}
	Port int
}

func getConfig() config {
	v := viper.New()
	v.AutomaticEnv()

	cfg := config{}
	cfg.Database.Host = v.GetString("DATABASE_HOST")
	cfg.Database.Port = v.GetInt("DATABASE_PORT")
	cfg.Database.Username = v.GetString("DATABASE_USERNAME")
	cfg.Database.Password = v.GetString("DATABASE_PASSWORD")
	cfg.Database.Database = v.GetString("DATABASE_DATABASE")
	cfg.Port = v.GetInt("PORT")

	return cfg
}

package database_test

import (
	"context"
	"testing"

	"github.com/devshark/alchemy-fleet/adapters/database"
	"github.com/devshark/alchemy-fleet/domain"
	"github.com/devshark/alchemy-fleet/ent"
	"github.com/devshark/alchemy-fleet/ent/enttest"
	"github.com/devshark/alchemy-fleet/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func newDatabase(t *testing.T) *ent.Client {
	t.Helper()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		enttest.WithOptions(ent.Debug()),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	return client
}

func TestGetSpacecrafts(t *testing.T) {
	// shadowing is intended
	require := require.New(t)
	t.Run("returns an empty list of spacecrafts", func(t *testing.T) {
		client := newDatabase(t)
		defer client.Close()

		ctx := context.Background()

		service := database.New(client)

		filters := domain.SpacecraftFilters{}

		spacecrafts, err := service.GetSpacecrafts(ctx, filters)
		require.NoError(err)
		require.Empty(spacecrafts)
	})
}

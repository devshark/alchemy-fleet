package database_test

import (
	"testing"

	"github.com/devshark/alchemy-fleet/ent"
	"github.com/devshark/alchemy-fleet/ent/enttest"
	"github.com/devshark/alchemy-fleet/ent/migrate"
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
	client := newDatabase(t)
	defer client.Close()

}

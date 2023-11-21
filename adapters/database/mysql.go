package database

import (
	"context"
	"fmt"

	"github.com/devshark/alchemy-fleet/adapters"
	"github.com/devshark/alchemy-fleet/domain"
	"github.com/devshark/alchemy-fleet/ent"
	"github.com/devshark/alchemy-fleet/ent/spacecraft"
)

type MySQLAdapter struct {
	client *ent.Client
}

// New creates a new MySQL adapter. It also creates the database schema.
// It will panic if it fails to create the schema. Stops as early as possible.
func New(client *ent.Client) adapters.Spacecraft {
	return &MySQLAdapter{
		client: client,
	}
}

// GetSpacecrafts returns a list of spacecrafts, with optional filters.
func (s *MySQLAdapter) GetSpacecrafts(ctx context.Context, filters domain.SpacecraftFilters) ([]*domain.Spacecraft, error) {
	// only return spacecrafts that are not deleted.
	query := s.client.Spacecraft.Query().Where(spacecraft.DeletedEQ(false))
	if filters.Name != "" {
		query = query.Where(spacecraft.NameContains(filters.Name))
	}

	if filters.Class != "" {
		query = query.Where(spacecraft.ClassContains(filters.Class))
	}

	// status must be an exact value, cannot be a wildcard.
	if filters.Status != domain.SpacecraftStatusUnknown {
		query = query.Where(spacecraft.StatusEQ(string(filters.Status)))
	}

	spacecrafts, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get spacecrafts: %w", err)
	}

	// create an empty slice with capacity equal to result length
	result := make([]*domain.Spacecraft, 0, len(spacecrafts))
	for _, s := range spacecrafts {
		spaceCraft := spacecraftToDomain(s)
		result = append(result, spaceCraft)
	}

	return result, nil
}

// GetSpacecraft returns a spacecraft by id. Throws an error if not found.
func (s *MySQLAdapter) GetSpacecraft(ctx context.Context, id int) (*domain.Spacecraft, error) {
	spacecraft, err := s.client.Spacecraft.Get(ctx, id)
	if ent.IsNotFound(err) {
		// for client's custom handling.
		return nil, fmt.Errorf("%w: %w", domain.ErrSpacecraftNotFound, err)
	}

	// anything else
	if err != nil {
		return nil, fmt.Errorf("failed to get spacecraft: %w", err)
	}

	// deleted is the same as not existing.
	if spacecraft.Deleted {
		return nil, domain.ErrSpacecraftNotFound
	}

	return spacecraftToDomain(spacecraft), nil
}

// CreateSpacecraft creates a new spacecraft, with database validation.
func (s *MySQLAdapter) CreateSpacecraft(ctx context.Context, spacecraft *domain.Spacecraft) (*domain.Spacecraft, error) {
	// create a new spacecraft
	newSpacecraft, err := s.client.Spacecraft.Create().
		SetName(spacecraft.Name).
		SetClass(spacecraft.Class).
		SetCrew(spacecraft.Crew).
		SetImage(spacecraft.Image).
		SetValue(spacecraft.Value).
		SetStatus(spacecraft.Status).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create spacecraft: %w", err)
	}

	return &domain.Spacecraft{
		Id:     newSpacecraft.ID,
		Name:   newSpacecraft.Name,
		Class:  newSpacecraft.Class,
		Crew:   newSpacecraft.Crew,
		Image:  newSpacecraft.Image,
		Value:  newSpacecraft.Value,
		Status: newSpacecraft.Status,
	}, nil
}

// UpdateSpacecraft updates an existing spacecraft, with database validation.
func (s *MySQLAdapter) UpdateSpacecraft(ctx context.Context, id int, spacecraft *domain.Spacecraft) (*domain.Spacecraft, error) {
	// get the spacecraft
	spacecraftToUpdate, err := s.client.Spacecraft.Get(ctx, id)
	if ent.IsNotFound(err) {
		// for client's custom handling.
		return nil, fmt.Errorf("%w: %w", domain.ErrSpacecraftNotFound, err)
	}

	// anything else
	if err != nil {
		return nil, fmt.Errorf("failed to get spacecraft: %w", err)
	}

	// deleted is the same as not existing.
	if spacecraftToUpdate.Deleted {
		return nil, domain.ErrSpacecraftNotFound
	}

	// update the spacecraft
	updatedSpacecraft, err := spacecraftToUpdate.Update().
		SetName(spacecraft.Name).
		SetClass(spacecraft.Class).
		SetCrew(spacecraft.Crew).
		SetImage(spacecraft.Image).
		SetValue(spacecraft.Value).
		SetStatus(spacecraft.Status).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update spacecraft: %w", err)
	}

	return &domain.Spacecraft{
		Id:     updatedSpacecraft.ID,
		Name:   updatedSpacecraft.Name,
		Class:  updatedSpacecraft.Class,
		Crew:   updatedSpacecraft.Crew,
		Image:  updatedSpacecraft.Image,
		Value:  updatedSpacecraft.Value,
		Status: updatedSpacecraft.Status,
	}, nil
}

// DeleteSpacecraft sets a spacecraft as deleted (soft delete). Throws an error if not found.
func (s *MySQLAdapter) DeleteSpacecraft(ctx context.Context, id int) error {
	// get the spacecraft
	spacecraftToDelete, err := s.client.Spacecraft.Get(ctx, id)
	if ent.IsNotFound(err) {
		// for client's custom handling.
		return fmt.Errorf("%w: %w", domain.ErrSpacecraftNotFound, err)
	}

	// anything else
	if err != nil {
		return fmt.Errorf("failed to get spacecraft: %w", err)
	}

	// deleted is the same as not existing.
	if spacecraftToDelete.Deleted {
		return domain.ErrSpacecraftNotFound
	}

	// set spacecraft as deleted
	err = spacecraftToDelete.Update().SetDeleted(true).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete spacecraft: %w", err)
	}

	return nil
}

package adapters

import (
	"context"

	"github.com/devshark/alchemy-fleet/domain"
)

// Spacecraft describes the methods that a Spacecraft adapter should implement.
type Spacecraft interface {
	GetSpacecrafts(ctx context.Context, filters domain.SpacecraftFilters) ([]*domain.Spacecraft, error)
	GetSpacecraft(ctx context.Context, id int) (*domain.Spacecraft, error)
	CreateSpacecraft(ctx context.Context, spacecraft *domain.Spacecraft) (*domain.Spacecraft, error)
	UpdateSpacecraft(ctx context.Context, id int, spacecraft *domain.Spacecraft) (*domain.Spacecraft, error)
	DeleteSpacecraft(ctx context.Context, id int) error
}

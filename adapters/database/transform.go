package database

import (
	"github.com/devshark/alchemy-fleet/domain"
	"github.com/devshark/alchemy-fleet/ent"
)

func spacecraftToDomain(spacecraft *ent.Spacecraft) *domain.Spacecraft {
	armaments := make([]domain.Armament, 0, len(spacecraft.Edges.Armament))
	// sometimes we don't need to query armament.
	if spacecraft.Edges.Armament != nil {
		for _, a := range spacecraft.Edges.Armament {
			armaments = append(armaments, domain.Armament{
				Title:    a.Title,
				Quantity: a.Quantity,
			})
		}
	}

	return &domain.Spacecraft{
		Id:       spacecraft.ID,
		Name:     spacecraft.Name,
		Class:    spacecraft.Class,
		Crew:     spacecraft.Crew,
		Image:    spacecraft.Image,
		Value:    spacecraft.Value,
		Status:   spacecraft.Status,
		Armament: armaments,
	}
}

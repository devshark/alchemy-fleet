package app

import (
	"encoding/json"
	"net/http"

	"github.com/devshark/alchemy-fleet/domain"
)

func (s *HttpServer) CreateSpacecraft(w http.ResponseWriter, r *http.Request) {
	var spacecraftRequest domain.Spacecraft

	err := json.NewDecoder(r.Body).Decode(&spacecraftRequest)
	if err != nil {
		// respond(w, http.StatusBadRequest, errorResponse{Error: "invalid request body"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	status := domain.StringToSpacecraftStatus(spacecraftRequest.Status)
	if status == domain.SpacecraftStatusUnknown {
		// respond(w, http.StatusBadRequest, errorResponse{Error: "invalid spacecraft status"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	ctx := r.Context()

	// spacecraft, err := s.spacecraftService.CreateSpacecraft(ctx, &spacecraftRequest)
	_, err = s.spacecraftService.CreateSpacecraft(ctx, &spacecraftRequest)
	if err != nil {
		// respond(w, http.StatusInternalServerError, errorResponse{Error: "failed to create spacecraft"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	// respond(w, http.StatusCreated, spacecraft)
	respond(w, http.StatusCreated, genericResponse{Success: true})
}

package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/devshark/alchemy-fleet/domain"
	"github.com/go-chi/chi/v5"
)

// UpdateSpacecraft updates a spacecraft by id, with request validation.
func (s *HttpServer) UpdateSpacecraft(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	if idParam == "" {
		// respond(w, http.StatusBadRequest, errorResponse{Error: "missing id parameter"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		// respond(w, http.StatusBadRequest, errorResponse{Error: "id must be a positive integer"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	var spacecraftRequest domain.Spacecraft

	err = json.NewDecoder(r.Body).Decode(&spacecraftRequest)
	if err != nil {
		// respond(w, http.StatusBadRequest, errorResponse{Error: "invalid request body"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	// if status is not empty, but the parsed value is unknown, then the status is invalid.
	status := domain.StringToSpacecraftStatus(spacecraftRequest.Status)
	if spacecraftRequest.Status != "" && status == domain.SpacecraftStatusUnknown {
		// respond(w, http.StatusBadRequest, errorResponse{Error: "invalid spacecraft status"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	ctx := r.Context()

	// spacecraft, err := s.spacecraftService.UpdateSpacecraft(ctx, id, &spacecraftRequest)
	_, err = s.spacecraftService.UpdateSpacecraft(ctx, id, &spacecraftRequest)
	if err != nil {
		// respond(w, http.StatusInternalServerError, errorResponse{Error: "failed to update spacecraft"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	// respond(w, http.StatusOK, spacecraft)
	respond(w, http.StatusOK, genericResponse{Success: true})
}

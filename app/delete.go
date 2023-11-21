package app

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/devshark/alchemy-fleet/domain"
	"github.com/go-chi/chi/v5"
)

func (s *HttpServer) DeleteSpacecraft(w http.ResponseWriter, r *http.Request) {
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

	ctx := r.Context()

	err = s.spacecraftService.DeleteSpacecraft(ctx, id)
	if errors.Is(err, domain.ErrSpacecraftNotFound) {
		// respond(w, http.StatusNotFound, errorResponse{Error: "spacecraft not found"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	// if found, but failed to delete, return an error
	if err != nil {
		// respond(w, http.StatusInternalServerError, errorResponse{Error: "failed to delete spacecraft"})
		respond(w, http.StatusBadRequest, genericResponse{Success: false})

		return
	}

	// respond(w, http.StatusOK, spacecraft)
	respond(w, http.StatusOK, genericResponse{Success: true})
}

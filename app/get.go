package app

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/devshark/alchemy-fleet/domain"
	"github.com/go-chi/chi/v5"
)

func (s *HttpServer) GetSpacecraftById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		respond(w, http.StatusBadRequest, errorResponse{Error: "missing id parameter"})

		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		respond(w, http.StatusBadRequest, errorResponse{Error: "invalid id parameter"})

		return
	}

	ctx := r.Context()

	spacecraft, err := s.spacecraftService.GetSpacecraft(ctx, id)
	if errors.Is(err, domain.ErrSpacecraftNotFound) {
		respond(w, http.StatusNotFound, errorResponse{Error: "spacecraft not found"})

		return
	}

	if err != nil {
		respond(w, http.StatusInternalServerError, errorResponse{Error: "failed to get spacecraft"})

		return
	}

	respond(w, http.StatusOK, spacecraft)
}

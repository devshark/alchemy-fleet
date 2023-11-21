package app

import (
	"net/http"
	"strings"

	"github.com/devshark/alchemy-fleet/domain"
)

// parse the request body into a domain.SpacecraftFilters struct
// returns a list of spacecrafts that match the filters
func (s *HttpServer) ListSpacecrafts(w http.ResponseWriter, r *http.Request) {
	filters := domain.SpacecraftFilters{}

	filters.Name = r.URL.Query().Get("name")
	filters.Class = r.URL.Query().Get("class")

	status := strings.TrimSpace(r.URL.Query().Get("status"))
	filters.Status = domain.StringToSpacecraftStatus(status)

	// if status is not empty, but the parsed value is unknown, then the status is invalid.
	if status != "" && filters.Status == domain.SpacecraftStatusUnknown {
		respond(w, http.StatusBadRequest, errorResponse{Error: "invalid status"})

		return
	}

	ctx := r.Context()

	spacecrafts, err := s.spacecraftService.GetSpacecrafts(ctx, filters)
	if err != nil {
		respond(w, http.StatusInternalServerError, errorResponse{Error: "failed to get spacecrafts"})

		return
	}

	// response schema as per documented requirement.
	data := struct {
		Data []*domain.Spacecraft `json:"data"`
	}{
		Data: spacecrafts,
	}

	respond(w, http.StatusOK, data)
}

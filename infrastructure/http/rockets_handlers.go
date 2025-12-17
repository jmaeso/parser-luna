package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jmaeso/parser-luna/app"
	"github.com/jmaeso/parser-luna/infrastructure/storage"
)

type RocketsHandler struct {
	RocketStateService app.RocketStateService
}

// ListRocketsHandler handles GET requests to /rockets endpoint.
// It returns a list of all known rocket states.
// The states are built at the time of calling the endpoint.
func (h *RocketsHandler) ListRockets(w http.ResponseWriter, r *http.Request) {
	domainRockets, err := h.RocketStateService.BuildAllRocketsState()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rockets := make([]Rocket, len(domainRockets))
	for i, dr := range domainRockets {
		rockets[i] = newRocketFromDomain(dr)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(rockets); err != nil {
		http.Error(w, "failed to marshal rocket state", http.StatusInternalServerError)
		return
	}
}

// GetRocketByID handles GET requests to /rockets/{id} endpoint.
// It returns the state of the given rocket.
// The state is built at the time of calling the endpoint.
func (h *RocketsHandler) GetRocketByID(w http.ResponseWriter, r *http.Request) {
	rocketID := r.PathValue("id")

	if rocketID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	domainRocket, err := h.RocketStateService.BuildRocketState(rocketID)
	if err != nil {
		if errors.Is(err, storage.ErrRocketNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpRocket := newRocketFromDomain(*domainRocket)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(httpRocket); err != nil {
		http.Error(w, "failed to marshal rocket state", http.StatusInternalServerError)
		return
	}
}

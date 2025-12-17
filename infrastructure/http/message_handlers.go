package http

import (
	"encoding/json"
	"net/http"

	"github.com/jmaeso/parser-luna/infrastructure/storage"
)

// MessageHandler is the responsible to handle the different method requests to /messages endpoint.
// Uses the storage to save them.
type MessageHandler struct {
	MessagesStorage storage.Messages
}

// HandlePostMessage handles POST requests to /messages endpoint.
// It stores them without any processing.
// It asumes valid payloads and does not validate them.
func (h *MessageHandler) PostMessage(w http.ResponseWriter, r *http.Request) {
	var payload PostMessagePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := payload.ToDomainMessage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := h.MessagesStorage.Insert(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

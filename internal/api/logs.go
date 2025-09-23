package api

import (
	"encoding/json"
	"net/http"

	"github.com/Observe86/intake-api/internal/model"
	"github.com/Observe86/intake-api/internal/service"
)

type LogsHandler struct {
	Service *service.LogsService
}

func (h *LogsHandler) HandleLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var logs []model.Log
	if err := json.NewDecoder(r.Body).Decode(&logs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.Process(logs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

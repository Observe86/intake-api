package api

import (
	"encoding/json"
	"net/http"

	"github.com/Observe86/intake-api/internal/model"
	"github.com/Observe86/intake-api/internal/service"
)

type TracesHandler struct {
	Service *service.TracesService
}

func (h *TracesHandler) HandleTraces(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var traces []model.Trace
	if err := json.NewDecoder(r.Body).Decode(&traces); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.Process(traces); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

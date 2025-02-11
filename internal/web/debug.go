package web

import (
	"encoding/json"
	"net/http"
)

func (h *handler) debugStore(w http.ResponseWriter, r *http.Request) {
	respone := make(map[string]interface{})
	respone["hosts"] = h.multiHostService.Hosts()
	containers, errors := h.multiHostService.ListAllContainers()
	respone["containers"] = containers
	respone["errors"] = errors

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respone)
}

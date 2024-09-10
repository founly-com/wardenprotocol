package http

import (
	"encoding/json"
	"net/http"

	"github.com/warden-protocol/wardenprotocol/prophet/api"
	"github.com/warden-protocol/wardenprotocol/prophet/internal/futures"
)

func (s *Server) postAck(w http.ResponseWriter, r *http.Request) {
	var req api.AckFuture
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ids := make([]futures.ID, len(req.IDs))
	for i, id := range req.IDs {
		ids[i] = futures.ID(id)
	}

	if err := s.sink.Ack(ids); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

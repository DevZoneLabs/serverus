package api

import (
	"encoding/json"
	"net/http"
)

type HealthCheckRequest struct {
	ChannelID string `json:"channel_id"`
}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {

	var hcr HealthCheckRequest
	if err := json.NewDecoder(r.Body).Decode(&hcr); err != nil {

		w.WriteHeader(http.StatusBadRequest)
	}

	err := s.bot.HealthCheckMessage(hcr.ChannelID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type SendMessageRequest struct {
	ChannelID string `json:"channel_id"`
	Message   string `json:"message"`
}

type SendMessageResponse struct {
	MessageId string `json:"message_id"`
}

func (s *Server) sendChannelMessage(w http.ResponseWriter, r *http.Request) {

	var smr SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&smr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg_id, err := s.bot.SendChannelMessage(smr.ChannelID, smr.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	out := SendMessageResponse{
		MessageId: *msg_id,
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

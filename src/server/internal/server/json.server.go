package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gsasso/go-backend/src/server/internal/ticker"
)

type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func ProvideAPIfunc(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

type JSONServer struct {
	listenAddr string
	svc        ticker.SummaryInt
}

func (s *JSONServer) Run() {

	http.ListenAndServe(s.listenAddr, nil)
}

func (s *JSONServer) HandleGetSummary(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	summary, err := s.svc.GetSummary()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(summary)
}

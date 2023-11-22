package server

import (
	"context"
	"net/http"
)

type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// func ProvideAPIfunc(fn APIFunc) http.HandlerFunc {

// }

// type JSONServer struct {
// 	listenAddr string
// }

package server

import (
	"net/http"

	"github.com/klimenkokayot/calc-api-go/http/server/handler"
	"github.com/klimenkokayot/calc-api-go/http/server/middleware"
)

func Run() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handler.CalcHandler)

	handler := middleware.RequestLogger(mux, "logger.txt")

	return handler
}

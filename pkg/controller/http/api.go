package api

import (
	"goProductExam/pkg/usecase"
	"log/slog"
	"net/http"
	"sync"
)

type api struct {
  mu *sync.Mutex
  r *http.ServeMux
  p usecase.ProductInterface
  log *slog.Logger
}

func New (mutex *sync.Mutex, router *http.ServeMux, p usecase.ProductInterface, log *slog.Logger) *api {
  return &api{mu: mutex, r: router, p: p, log: log}
}

func (api *api) Handle() {
  api.r.HandleFunc("/api/v1/product", api.product)
}

func (api *api) ListenAndServe(addres string) error {
  return http.ListenAndServe(addres, api.r)
}

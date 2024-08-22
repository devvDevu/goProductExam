package api

import (
	"goProductExam/pkg/repository"
	"log/slog"
	"net/http"
	"sync"
)

type api struct {
  mu *sync.Mutex
  r *http.ServeMux
  db *repository.PGRepo
  log *slog.Logger
}

func New (mutex *sync.Mutex, router *http.ServeMux, db *repository.PGRepo, log *slog.Logger) *api {
  return &api{mu: mutex, r: router, db: db, log: log}
}

func (api *api) Handle() {
  api.r.HandleFunc("api/v1/get_products", api.product)
}

func (api *api) ListenAndServe(addres string) error {
  return http.ListenAndServe(addres, api.r)
}

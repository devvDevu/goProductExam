package main

import (
	"fmt"
	"goProductExam/internal/config"
  "goProductExam/pkg/infrastructure/repository"
  "goProductExam/pkg/controller/http"
  "goProductExam/pkg/usecase"
	"log/slog"
	"net/http"
	"os"
	"sync"
)

const (
  envLocal = "local"
  envDev = "dev"
  envProd = "prod"
)

func main() {
  cfg := config.MustLoad()
  log := setupLogger(cfg.Env)
  p_db, err := repository.SetupPGRepo(cfg.StoragePath)
  productRepo := usecase.New(p_db)
  if err != nil {
    fmt.Printf("log: %v\n", err)
    os.Exit(1)
  }
  api := api.New(&sync.Mutex{},&http.ServeMux{}, productRepo, log)
  api.Handle()
  err = api.ListenAndServe(cfg.HttpServer.Address)
  if err != nil {
    fmt.Printf("log: %v\n", err)
    os.Exit(1)
  }

}

func setupLogger(env string) *slog.Logger {
  var log *slog.Logger
  switch env{
    case "local":
      log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
    case "dev":
      log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
    case "prod":
      log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
  }

  return log
}

package repository

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGRepo struct {
  mu sync.Mutex
  pool *pgxpool.Pool
}

func SetupPGRepo (storagePath string) (*PGRepo, error) {
  ctx := context.Background()
  pool, err := pgxpool.New(ctx, storagePath)
  if err != nil {
    log.Printf("Connection failed, storagePath: %s", storagePath)
    return nil, err
  }

  _, err = pool.Exec(ctx, `
  create table if not exists product(
    id serial primary key,
    name varchar not null,
    cost integer not null default 0
  );
  `)
  if err != nil {
    log.Printf("Failed to create table")
    return nil, err
  }
  return &PGRepo{mu: sync.Mutex{}, pool: pool}, nil
}

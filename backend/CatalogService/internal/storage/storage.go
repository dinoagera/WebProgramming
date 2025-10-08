package storage

import (
	"catalogservice/internal/models"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	log  *slog.Logger
	Pool *pgxpool.Pool
}

func New(log *slog.Logger, storgaePath string) (*Storage, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, storgaePath)
	if err != nil {
		log.Info("failed to connect to pgxpool", "err", err)
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Info("failed to ping to db", "err", err)
		return nil, err
	}
	return &Storage{
		log:  log,
		Pool: pool,
	}, nil
}
func (s *Storage) GetCatalog() (models.Goods, error) {

}

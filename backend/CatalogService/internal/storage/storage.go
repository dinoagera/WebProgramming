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
func (s *Storage) GetCatalog() ([]models.Good, error) {
	rows, err := s.Pool.Query(context.Background(), `SELECT product_id,category, sex, size, price, color, tag FROM Goods`)
	if err != nil {
		s.log.Info("failed to get catalog query", "err", err)
		return nil, err
	}
	var goods []models.Good
	defer rows.Close()
	for rows.Next() {
		var good models.Good
		err := rows.Scan(
			&good.ProductID,
			&good.Category,
			&good.Sex,
			&good.Size,
			&good.Price,
			&good.Color,
			&good.Tag,
		)
		if err != nil {
			s.log.Info("failed to scan to struct", "err", err)
			return nil, err
		}
		goods = append(goods, good)
	}
	if err := rows.Err(); err != nil {
		s.log.Info("failed to reading rows", "err", err)
		return nil, err
	}
	return goods, nil
}

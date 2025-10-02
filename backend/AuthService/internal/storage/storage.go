package storage

import (
	"authservice/internal/domain/models"
	"authservice/internal/handler"
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	log  *slog.Logger
	Pool *pgxpool.Pool
}

func New(log *slog.Logger, storagePath string) (*Storage, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, storagePath)
	if err != nil {
		log.Info("failed to connect to pgxpool", "err:", err)
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Info("failed to to ping to db", "err", err)
		return nil, err
	}
	return &Storage{
		log:  log,
		Pool: pool,
	}, nil
}
func (s *Storage) CreateUser(email, passHash string) error {
	var exists bool
	err := s.Pool.QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`, email).Scan(&exists)
	if err != nil {
		s.log.Info("failed to query row", "err", err)
		return err
	}
	if exists {
		s.log.Info("user is busy", "email", email)
		return handler.ErrEmailBusy
	}
	_, err = s.Pool.Exec(context.Background(), `INSERT INTO users (email, pass_hash, created_at) VALUES ($1,$2,$3)`, email, passHash, time.Now())
	if err != nil {
		s.log.Info("failed to execute query on create user", "err:", err)
		return err
	}
	s.log.Info("user created", "email", email)
	return nil
}
func (s *Storage) LoginUser(email string) (models.User, error) {
	var user models.User
	err := s.Pool.QueryRow(context.Background(), `SELECT uid, email, pass_hash FROM users WHERE email=$1`, email).Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		s.log.Info("failed to execute query on login user", "err:", err)
		return user, err
	}
	return user, nil
}

package storage

import (
	"catalogservice/internal/models"
	"catalogservice/lib"
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	log           *slog.Logger
	Pool          *pgxpool.Pool
	imageBasePath string
}

func New(log *slog.Logger, storgaePath string, imageBasePath string) (*Storage, error) {
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
		log:           log,
		Pool:          pool,
		imageBasePath: imageBasePath,
	}, nil
}

func (s *Storage) GetCatalog() ([]models.Good, error) {
	rows, err := s.Pool.Query(context.Background(), `SELECT product_id,category, sex, sizes, price, color, tag, '/api/image/' || product_id as image_url FROM goods`)
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
			&good.Sizes,
			&good.Price,
			&good.Color,
			&good.Tag,
			&good.ImageURL,
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
	if len(goods) == 0 {
		return []models.Good{}, lib.ErrCatalogIsEmpty
	}
	return goods, nil
}
func (s *Storage) GetImage(productID string) ([]byte, error) {
	numericID := filepath.Base(productID)
	filename := fmt.Sprintf("productID%s.png", numericID)
	imagePath := filepath.Join(s.imageBasePath, filename)
	files, err := filepath.Glob(filepath.Join(s.imageBasePath, "productID*.png"))
	if err != nil {
		s.log.Info("DEBUG: Failed to list files", "err", err)
	} else {
		s.log.Info("DEBUG: Available product images", "files", files)
	}
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		s.log.Info("PNG image not found in filesystem",
			"product_id", productID,
			"path", imagePath,
			"file_exists", false)
		return nil, lib.ErrImageNotFound
	} else {
		s.log.Info("DEBUG: File exists!", "path", imagePath)
	}
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		s.log.Info("failed to read PNG file", "err", err, "path", imagePath)
		return nil, err
	}
	s.log.Info("successfully loaded PNG image from filesystem", "product_id", productID, "size", len(imageData))
	return imageData, nil
}
func (s *Storage) GetFavourites(userID int) ([]models.Favourites, error) {
	rows, err := s.Pool.Query(context.Background(),
		`SELECT 
		g.product_id,
		g.category,
		g.sex,
		g.sizes, 
		g.price, 
		g.color, 
		g.tag, '/api/image/' || g.product_id as image_url 
		FROM favourites f
		JOIN goods g ON f.product_id = g.product_id
		JOIN users u ON f.uid = u.uid
		WHERE u.uid = $1 `, userID)
	if err != nil {
		s.log.Info("failed to get catalog query", "err", err)
		return nil, err
	}
	var favourites []models.Favourites
	defer rows.Close()
	for rows.Next() {
		var good models.Favourites
		err := rows.Scan(
			&good.ProductID,
			&good.Category,
			&good.Sex,
			&good.Sizes,
			&good.Price,
			&good.Color,
			&good.Tag,
			&good.ImageURL,
		)
		if err != nil {
			s.log.Info("failed to scan to struct", "err", err)
			return nil, err
		}
		favourites = append(favourites, good)
	}
	if err := rows.Err(); err != nil {
		s.log.Info("failed to reading rows", "err", err)
		return nil, err
	}
	if len(favourites) == 0 {
		return []models.Favourites{}, lib.ErrFavouritesIsEmpty
	}
	return favourites, nil
}

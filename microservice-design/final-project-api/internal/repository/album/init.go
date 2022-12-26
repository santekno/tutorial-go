package repository

import (
	"context"
	"database/sql"

	"final-project-api/internal/entity"
	"final-project-api/internal/repository/album/cache"
	"final-project-api/internal/repository/album/psql"

	"github.com/go-redis/redis/v8"
)

type AlbumRepository interface {
	Get(ctx context.Context, id int64) (*entity.Album, error)
	Create(ctx context.Context, album *entity.Album) (int64, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	BatchCreate(ctx context.Context, albums []entity.Album) ([]int64, error)
	Update(ctx context.Context, album entity.Album) error
	Delete(ctx context.Context, id int64) error

	GetAlbumCache(ctx context.Context, id int64) (*entity.Album, error)
	GetAllAlbumCache(ctx context.Context) ([]entity.Album, error)
	SetAlbumCache(ctx context.Context, id int64, album entity.Album) error
	SetAllAlbumCache(ctx context.Context, albums []entity.Album) error
	DeleteAlbumCache(ctx context.Context, id int64) error
}

type albumRepository struct {
	postgres psql.AlbumPostgres
	cache    cache.AlbumPostgres
}

// The function is to initialize the album repository
func NewAlbumRepository(db *sql.DB, client *redis.Client) AlbumRepository {
	return &albumRepository{
		postgres: psql.NewAlbumPostgres(db),
		cache:    cache.NewAlbumRedis(client),
	}
}

package cache

import (
	"context"

	"final-project-api/internal/entity"

	"github.com/go-redis/redis/v8"
)

type AlbumPostgres interface {
	GetAlbum(ctx context.Context, id int64) (*entity.Album, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	SetAlbum(ctx context.Context, id int64, album entity.Album) error
	SetAllAlbum(ctx context.Context, albums []entity.Album) error
	Delete(ctx context.Context, id int64) error
}

type albumConnection struct {
	client *redis.Client
}

// The function is to initialize the album psql repository
func NewAlbumRedis(cache *redis.Client) AlbumPostgres {
	return &albumConnection{client: cache}
}

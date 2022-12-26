package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"final-project-api/internal/entity"

	"github.com/go-redis/redis/v8"
)

// Get Specific album cache
func (repo *albumConnection) GetAlbum(ctx context.Context, id int64) (*entity.Album, error) {
	var album entity.Album

	key := fmt.Sprintf(albumDetailKey, id)

	albumsString, err := repo.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return &album, nil
	}
	if err != nil {
		return &album, err
	}

	err = json.Unmarshal([]byte(albumsString), &album)
	if err != nil {
		return &album, err
	}

	return &album, nil
}

// GetAllAlbum is function to get all albums from database
func (repo *albumConnection) GetAllAlbum(ctx context.Context) ([]entity.Album, error) {
	var albums []entity.Album

	albumsString, err := repo.client.Get(ctx, albumsKey).Result()
	if err == redis.Nil {
		return albums, nil
	}
	if err != nil {
		return albums, err
	}

	err = json.Unmarshal([]byte(albumsString), &albums)
	if err != nil {
		return albums, err
	}

	return albums, nil
}

func (repo *albumConnection) SetAlbum(ctx context.Context, id int64, album entity.Album) error {
	key := fmt.Sprintf(albumDetailKey, id)

	albumsString, err := json.Marshal(album)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, key, albumsString, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *albumConnection) SetAllAlbum(ctx context.Context, albums []entity.Album) error {
	albumsString, err := json.Marshal(albums)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, albumsKey, albumsString, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *albumConnection) Delete(ctx context.Context, id int64) error {
	key := fmt.Sprintf(albumDetailKey, id)

	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}

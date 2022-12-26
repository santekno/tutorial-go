package usecase

import (
	"context"

	albumRepository "final-project-api/internal/repository/album"

	"final-project-api/internal/entity"
)

type AlbumUsecase interface {
	Get(ctx context.Context, id int64) (*entity.Album, error)
	Create(ctx context.Context, album *entity.Album) (*entity.Album, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	BatchCreate(ctx context.Context, albums []entity.Album) ([]entity.Album, error)
	Update(ctx context.Context, album entity.Album) (entity.Album, error)
	Delete(ctx context.Context, id int64) error
}

type albumUsecase struct {
	albumRepository albumRepository.AlbumRepository
}

// The function is to initialize the album usecase
func NewAlbumUsecase(albumRepository albumRepository.AlbumRepository) AlbumUsecase {
	return &albumUsecase{
		albumRepository: albumRepository,
	}
}

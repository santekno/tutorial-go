package artist

import (
	"context"

	"final-project-api/internal/entity"
	artistRepository "final-project-api/internal/repository/artist"
)

type ArtistUsecase interface {
	Get(ctx context.Context, id int64) (*entity.Artist, error)
	GetAll(ctx context.Context, limit, offset int32) (*[]entity.Artist, error)
}

type artistUsecase struct {
	artistRepository artistRepository.ArtistRepository
}

// The function is to initialize the artist usecase
func NewArtistUsecase(artistRepository artistRepository.ArtistRepository) ArtistUsecase {
	return &artistUsecase{
		artistRepository: artistRepository,
	}
}

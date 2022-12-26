package artist

import (
	"context"

	"final-project-api/internal/entity"
)

func (usecase *artistUsecase) Get(ctx context.Context, id int64) (*entity.Artist, error) {
	artist, err := usecase.artistRepository.GetById(ctx, id)
	if err != nil {
		return artist, err
	}

	if artist.ID != 0 {
		return artist, nil
	}

	return artist, nil
}

func (usecase *artistUsecase) GetAll(ctx context.Context, limit, offset int32) (*[]entity.Artist, error) {
	artists, err := usecase.artistRepository.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

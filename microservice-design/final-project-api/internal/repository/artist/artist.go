package repository

import (
	"context"

	"final-project-api/internal/entity"
)

// It will call the function get all in grpc/artist
func (repo *artistRepository) GetAll(ctx context.Context, limit, offset int32) (*[]entity.Artist, error) {
	return repo.grpc.GetAll(ctx, limit, offset)
}

// It will call the function get by id in psql/artist
func (repo *artistRepository) GetById(ctx context.Context, id int64) (*entity.Artist, error) {
	return repo.grpc.GetById(ctx, id)
}

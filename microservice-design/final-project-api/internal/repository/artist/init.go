package repository

import (
	"context"

	"final-project-api/internal/entity"
	"final-project-api/internal/repository/artist/grpc"
)

type ArtistRepository interface {
	GetAll(ctx context.Context, limit, offset int32) (*[]entity.Artist, error)
	GetById(ctx context.Context, id int64) (*entity.Artist, error)
}

type artistRepository struct {
	grpc grpc.ArtistGrpc
}

// The function is to initialize the artist repository
func NewArtistRepository(client grpc.ArtistInterface) ArtistRepository {
	return &artistRepository{
		grpc: grpc.NewArtistGrpc(client),
	}
}

package grpc

import (
	"context"

	"final-project-api/internal/entity"
	"final-project-api/internal/proto"
)

// GetAllartist is function to get all artists from grpc call
func (repo *artistGrpc) GetAll(ctx context.Context, limit, offset int32) (*[]entity.Artist, error) {
	resp, err := repo.grpc.GetAll(ctx, &proto.PagingRequest{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var artists = make([]entity.Artist, 0)
	for _, artist := range resp.GetArtists() {
		artists = append(artists, entity.Artist{
			ID:     int64(artist.GetId()),
			Name:   artist.GetName(),
			Rating: int(artist.GetRating()),
		})
	}

	return &artists, nil
}

// Get is function to get specific artist by id from grpc call
func (repo *artistGrpc) GetById(ctx context.Context, id int64) (*entity.Artist, error) {
	resp, err := repo.grpc.GetById(ctx, &proto.ArtistRequest{
		Id: int32(id),
	})
	if err != nil {
		return nil, err
	}

	artist := &entity.Artist{}
	artist.ID = int64(resp.GetId())
	artist.Name = resp.GetName()
	artist.Rating = int(resp.GetRating())
	return artist, nil
}

package grpc

import (
	"context"

	"final-project-api/internal/entity"
	"final-project-api/internal/proto"

	"google.golang.org/grpc"
)

type ArtistGrpc interface {
	GetAll(ctx context.Context, limit, offset int32) (*[]entity.Artist, error)
	GetById(ctx context.Context, id int64) (*entity.Artist, error)
}

type ArtistInterface interface {
	GetById(ctx context.Context, in *proto.ArtistRequest, opts ...grpc.CallOption) (*proto.ArtistResponse, error)
	GetAll(ctx context.Context, in *proto.PagingRequest, opts ...grpc.CallOption) (*proto.ArtistsResponse, error)
}

type artistGrpc struct {
	grpc proto.ArtistClient
}

// The function is to initialize the artist grpc repository
func NewArtistGrpc(client ArtistInterface) ArtistGrpc {
	return &artistGrpc{grpc: client}
}

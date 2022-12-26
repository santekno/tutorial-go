package config

import (
	"log"

	proto "final-project-api/internal/proto"

	"google.golang.org/grpc"
)

// open connection into grpc artist service
func ServiceGrpcArtist(grpcUrl string) proto.ArtistClient {
	conn, err := grpc.Dial(grpcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", grpcUrl, err)
	}

	return proto.NewArtistClient(conn)
}

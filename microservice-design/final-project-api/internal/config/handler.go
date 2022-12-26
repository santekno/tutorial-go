package config

import (
	albumHandler "final-project-api/internal/handler/album"
	artistHandler "final-project-api/internal/handler/artist"
	albumUsecase "final-project-api/internal/usecase/album"
	artistUsecase "final-project-api/internal/usecase/artist"
)

type Handler struct {
	AlbumHandler  albumHandler.AlbumHandler
	ArtistHandler artistHandler.ArtistHandler
}

// Function to initialize handler
func InitHandler(albumUsecase albumUsecase.AlbumUsecase, artistUsecase artistUsecase.ArtistUsecase) Handler {
	return Handler{
		AlbumHandler:  albumHandler.NewAlbumHandler(albumUsecase),
		ArtistHandler: artistHandler.NewArtistHandler(artistUsecase),
	}
}

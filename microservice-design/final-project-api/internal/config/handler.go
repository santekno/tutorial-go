package config

import (
	albumHandler "final-project-api/internal/handler/album"
	albumUsecase "final-project-api/internal/usecase/album"
)

type Handler struct {
	AlbumHandler albumHandler.AlbumHandler
}

// Function to initialize handler
func InitHandler(albumUsecase albumUsecase.AlbumUsecase) Handler {
	return Handler{
		AlbumHandler: albumHandler.NewAlbumHandler(albumUsecase),
	}
}

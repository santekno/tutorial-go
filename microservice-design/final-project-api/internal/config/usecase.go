package config

import (
	albumRepository "final-project-api/internal/repository/album"
	albumUsecase "final-project-api/internal/usecase/album"
)

type Usecase struct {
	AlbumUsecase albumUsecase.AlbumUsecase
}

// Function to initialize usecase
func InitUsecase(albumRepository albumRepository.AlbumRepository) Usecase {
	return Usecase{
		AlbumUsecase: albumUsecase.NewAlbumUsecase(albumRepository),
	}
}

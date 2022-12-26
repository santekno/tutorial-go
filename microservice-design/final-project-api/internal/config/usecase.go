package config

import (
	albumRepository "final-project-api/internal/repository/album"
	artistRepository "final-project-api/internal/repository/artist"
	albumUsecase "final-project-api/internal/usecase/album"
	artistUsecase "final-project-api/internal/usecase/artist"
)

type Usecase struct {
	AlbumUsecase  albumUsecase.AlbumUsecase
	ArtistUsecase artistUsecase.ArtistUsecase
}

// Function to initialize usecase
func InitUsecase(albumRepository albumRepository.AlbumRepository, artistRepository artistRepository.ArtistRepository) Usecase {
	return Usecase{
		AlbumUsecase:  albumUsecase.NewAlbumUsecase(albumRepository),
		ArtistUsecase: artistUsecase.NewArtistUsecase(artistRepository),
	}
}

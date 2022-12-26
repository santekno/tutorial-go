package artist

import (
	artistUsecase "final-project-api/internal/usecase/artist"

	"github.com/gin-gonic/gin"
)

type ArtistHandler interface {
	Get(context *gin.Context)
	GetAll(context *gin.Context)
}

type artistHandler struct {
	artistUsecase artistUsecase.ArtistUsecase
}

// The function is to initialize the artist handler
func NewArtistHandler(artistUsecase artistUsecase.ArtistUsecase) ArtistHandler {
	return &artistHandler{
		artistUsecase: artistUsecase,
	}
}

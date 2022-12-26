package album

import (
	albumUsecase "final-project-api/internal/usecase/album"

	"github.com/gin-gonic/gin"
)

type AlbumHandler interface {
	Get(context *gin.Context)
	Create(context *gin.Context)
	GetAllAlbum(context *gin.Context)
	BatchCreate(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type albumHandler struct {
	albumUsecase albumUsecase.AlbumUsecase
}

// The function is to initialize the album handler
func NewAlbumHandler(albumUsecase albumUsecase.AlbumUsecase) AlbumHandler {
	return &albumHandler{
		albumUsecase: albumUsecase,
	}
}

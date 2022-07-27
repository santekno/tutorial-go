package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/santekno/postgres-go/services"
)

type AlbumHandler struct {
	AlbumService services.AlbumService
}

func NewHandler(service services.AlbumService) *AlbumHandler {
	return &AlbumHandler{
		AlbumService: service,
	}
}

func (ah *AlbumHandler) GetAllAlbums(c *gin.Context) {
	albums, err := ah.AlbumService.GetAllAlbum()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error when find get all album",
		})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func (ah *AlbumHandler) Get(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, map[string]interface{}{
			"message": "invalid exercise id",
		})
		return
	}

	album, err := ah.AlbumService.Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error when find get album by id : " + paramID,
		})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (ah *AlbumHandler) GetByArtist(c *gin.Context) {
	paramArtist := c.Request.FormValue("artist")
	if paramArtist == "" {
		c.JSON(400, map[string]interface{}{
			"message": "invalid artist",
		})
		return
	}

	album, err := ah.AlbumService.GetByArtist(paramArtist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error when find get album by artist : " + paramArtist,
		})
		return
	}

	c.JSON(http.StatusOK, album)
}

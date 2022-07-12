package album

import (
	"github.com/santekno/using-gorm/internal/domain"
)

type AlbumRepo interface {
	GetAllAlbum() ([]domain.Album, error)
	GetAlbumById(id int64) (domain.Album, error)
	CreateAlbum(album domain.Album) error
	BulkInsertAlbum(albums []domain.Album) error
	UpdateAlbum(album domain.Album) error
	DeleteAlbum(id int) error
}

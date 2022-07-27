package services

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

type AlbumService interface {
	Get(id int64) (*Album, error)
	Create(album *Album) error
	GetAllAlbum() ([]Album, error)
	BatchCreate(albums []Album) error
	Update(Album Album) error
	Delete(id int64) error
}

package services

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

type AlbumService interface {
	Get(id int64) (*Album, error)
	Create(album *Album) error
	GetAllAlbum() ([]Album, error)
	BatchCreate(albums []Album) error
	Update(Album Album) error
	Delete(id int64) error
	GetByArtist(artist string) ([]Album, error)
}

package repository

import (
	"errors"
	"log"

	"github.com/santekno/using-gorm/internal/domain"
	"gorm.io/gorm"
)

type AlbumRepo struct {
	db *gorm.DB
}

func NewAlbumRepo(database *gorm.DB) *AlbumRepo {
	return &AlbumRepo{
		db: database,
	}
}

func (al *AlbumRepo) BulkInsertAlbum(albums []domain.Album) error {
	err := al.db.CreateInBatches(albums, len(albums))
	if err.Error != nil {
		log.Printf("error when bulk insert err: %v", err)
		return err.Error
	}

	log.Print("Success bulk insert Album")
	return nil
}

func (al *AlbumRepo) GetAllAlbum() ([]domain.Album, error) {
	var album []domain.Album
	err := al.db.Find(&album).Error
	if err != nil {
		log.Printf("error when getting album err: %v", err)
		return nil, err
	}
	return album, nil
}

func (al *AlbumRepo) GetAlbumById(id int64) (domain.Album, error) {
	var album domain.Album
	err := al.db.Where("id = ?", id).Take(&album).Error
	if err != nil {
		log.Printf("error when getting album err: %v", err)
		return domain.Album{}, err
	}
	return album, nil
}

func (al *AlbumRepo) CreateAlbum(album domain.Album) error {
	tx := al.db.Create(album)
	if tx.Error != nil {
		return tx.Error
	}
	log.Print("Success create Album")
	return nil
}

func (al *AlbumRepo) UpdateAlbum(album domain.Album) error {
	var albumUpdate domain.Album
	al.db.Where("id = ?", album.ID).Find(&albumUpdate)
	if albumUpdate.ID == 0 {
		return errors.New("not found")
	}

	albumUpdate.Title = album.Title
	albumUpdate.Artist = album.Artist
	albumUpdate.Price = album.Price

	err := al.db.Save(albumUpdate).Error
	if err != nil {
		return err
	}
	log.Print("Success update Album")
	return nil
}

func (al *AlbumRepo) DeleteAlbum(id int) error {
	var albumUpdate domain.Album
	al.db.Where("id = ?", id).Find(&albumUpdate)
	al.db.Delete(albumUpdate)
	log.Print("Success delete Album")
	return nil
}

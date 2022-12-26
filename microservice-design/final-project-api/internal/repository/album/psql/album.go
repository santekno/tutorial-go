package psql

import (
	"context"
	"fmt"
	"log"
	"time"

	"final-project-api/internal/entity"
)

// Create is function to create album to database
func (repo *albumConnection) Create(ctx context.Context, album *entity.Album) (int64, error) {
	// The query insert
	query := `
        INSERT INTO public.album (title, artist, price) 
        VALUES ($1, $2, $3)
        RETURNING id`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query, album.Title, album.Artist, album.Price).Scan(&album.ID)
	if err != nil {
		return 0, err
	}

	return album.ID, nil
}

// Get is function to get specific album by id from database
func (repo *albumConnection) Get(ctx context.Context, id int64) (*entity.Album, error) {
	// The query select
	query := `
        SELECT id, title, artist, price
        FROM album
        WHERE id = $1`

	var album entity.Album

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query and find the specific album and then set the result to album variable
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	// If any error
	if err != nil {
		return nil, err
	}

	return &album, nil
}

// GetAllAlbum is function to get all albums from database
func (repo *albumConnection) GetAllAlbum(ctx context.Context) ([]entity.Album, error) {
	// The query select
	query := `
		SELECT id, artist, title, price
		FROM album`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var albums []entity.Album

	// Run the query to get all albums
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// If the albums is not empty
	for rows.Next() {
		var album entity.Album

		// Set to the album variable
		err := rows.Scan(
			&album.ID,
			&album.Artist,
			&album.Title,
			&album.Price,
		)
		// If any error
		if err != nil {
			return nil, err
		}

		// add the album to albums variable
		albums = append(albums, album)
	}

	return albums, nil
}

// BatchCreate is function to insert some albums in once to database
func (repo *albumConnection) BatchCreate(ctx context.Context, albums []entity.Album) ([]int64, error) {
	var IDs []int64

	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return IDs, nil
	}
	// If any error, the transaction will be rollback
	defer tx.Rollback()

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query insert
	query := `INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id`

	// Loop every album
	for _, album := range albums {
		var id int64

		// Run query insert of every album to database
		err := tx.QueryRowContext(ctx, query, album.Title, album.Artist, album.Price).Scan(&id)
		if err != nil {
			log.Printf("error execute insert err: %v", err)
			continue
		}

		// Add the new id to IDs variable
		IDs = append(IDs, id)
	}

	// Commit the transaction
	err = tx.Commit()
	// If any error
	if err != nil {
		return IDs, err
	}

	return IDs, nil
}

// Update is function to update album in database
func (repo *albumConnection) Update(ctx context.Context, album entity.Album) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query update
	query := `UPDATE album set title=$1, artist=$2, price=$3 WHERE id=$4`

	// Run the query
	result, err := repo.db.ExecContext(ctx, query, album.Title, album.Artist, album.Price, album.ID)
	if err != nil {
		return err
	}

	// Get how many data has been updated
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected update : %d", rows)
	return nil
}

// Delete is function to delete album in database
func (repo *albumConnection) Delete(ctx context.Context, id int64) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query delete
	query := `DELETE from album WHERE id=$1`

	// Run the delete query
	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Get how many data has been deleted
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected delete : %d", rows)
	return nil
}

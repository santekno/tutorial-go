package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type PostgresService struct {
	db *sql.DB
}

func NewPostgresService(db *sql.DB) *PostgresService {
	return &PostgresService{db: db}
}

func (p *PostgresService) Create(album *Album) error {
	query := `
        INSERT INTO public.album (id, title, artist, price) 
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return p.db.QueryRowContext(ctx, query, album.Title, album.Artist, album.Price).Scan(&album.ID)
}

func (p *PostgresService) Get(id int64) (*Album, error) {
	query := `
        SELECT id, title, artist, price
        FROM public.album
        WHERE id = $1`

	var album Album

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	if err != nil {
		return nil, err
	}

	return &album, nil
}

func (p *PostgresService) GetAllAlbum() ([]Album, error) {
	query := `
		SELECT id, title, artist, price
		FROM album`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var albums []Album

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album Album
		err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Artist,
			&album.Price,
		)
		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func (p *PostgresService) BatchCreate(albums []Album) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)`

	for _, album := range albums {
		_, err := tx.ExecContext(ctx, query, album.Title, album.Artist, album.Price)
		if err != nil {
			log.Printf("error execute insert err: %v", err)
			continue
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresService) Update(album Album) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `UPDATE album set title=$1, artist=$2, price=$3 WHERE id=$4`
	result, err := p.db.ExecContext(ctx, query, album.Title, album.Artist, album.Price, album.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected update : %d", rows)
	return nil
}

func (p *PostgresService) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `DELETE from album WHERE id=$1`
	result, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Affected delete : %d", rows)
	return nil
}

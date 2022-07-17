package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type MSSQLService struct {
	db *sql.DB
}

func NewMSSQLService(db *sql.DB) *MSSQLService {
	return &MSSQLService{db: db}
}

func (p *MSSQLService) Create(album *Album) error {
	query := `
        INSERT INTO dbo.album (title, artist, price) 
        VALUES (@title, @artist, @price)`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := p.db.ExecContext(ctx, query,
		sql.Named("title", album.Title),
		sql.Named("artist", album.Artist),
		sql.Named("price", album.Price))
	if err != nil {
		return err
	}

	return nil
}

func (p *MSSQLService) Get(id int64) (*Album, error) {
	query := `
        SELECT id, title, artist, price
        FROM dbo.album
        WHERE id = @id`

	var album Album

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := p.db.QueryRowContext(ctx, query, sql.Named("id", id)).Scan(
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

func (p *MSSQLService) GetAllAlbum() ([]Album, error) {
	query := `
		SELECT id, title, artist, price
		FROM dbo.album`

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

func (p *MSSQLService) BatchCreate(albums []Album) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `INSERT INTO dbo.album (title, artist, price) VALUES (@title, @artist, @price)`

	for _, album := range albums {
		_, err := tx.ExecContext(ctx, query,
			sql.Named("title", album.Title),
			sql.Named("artist", album.Artist),
			sql.Named("price", album.Price))
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

func (p *MSSQLService) Update(album Album) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `UPDATE dbo.album set title=@title, artist=@artist, price=@price WHERE id=@id`
	result, err := p.db.ExecContext(ctx, query,
		sql.Named("title", album.Title),
		sql.Named("artist", album.Artist),
		sql.Named("price", album.Price),
		sql.Named("id", album.ID))
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

func (p *MSSQLService) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `DELETE from dbo.album WHERE id=@id`
	result, err := p.db.ExecContext(ctx, query, sql.Named("id", id))
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

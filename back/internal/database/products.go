package database

import (
	"context"
)

type Product struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (db *DB) InsertProduct(name string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO products (name)
		VALUES ($1)`

	result, err := db.ExecContext(ctx, query, name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

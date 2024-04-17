package database

import "context"

type Color struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Code string `db:"code"`
}

func (db *DB) InsertColor(name, code string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO colors (name, code)
		VALUES ($1, $2)`

	result, err := db.ExecContext(ctx, query, name, code)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

package database

import "context"

type Status struct {
	id   int    `db:"id"`
	name string `db:"name"`
	code string `db:"code"`
}

func (db *DB) InsertStatus(name, code string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO statuses (name, code)
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

package database

import (
	"context"
)

type Product struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Image    string `db:"image" json:"image"`
	Price    string `db:"price" json:"price"`
	ColorId  string `db:"color_id" json:"color_id"`
	StatusId string `db:"status_id" json:"status_id"`
}

func (db *DB) InsertProduct(name, image string, price, color_id, status_id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO products (name, image, price, color_id, status_id)
		VALUES ($1, $2, $3, $4, $5)`

	result, err := db.ExecContext(ctx, query, name, image, price, color_id, status_id)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}
func (db *DB) GetProducts() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	list := []Product{}

	query := `SELECT * FROM products`

	err := db.SelectContext(ctx, &list, query)

	if err != nil {
		return list, nil
	}

	return list, err
}

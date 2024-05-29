package database

import (
	"context"
)

type Product struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Image    string `db:"image" json:"image"`
	Price    string `db:"price" json:"price"`
	ColorId  int    `db:"color_id" json:"color_id"`
	StatusId int    `db:"status_id" json:"status_id"`
	Color    Color  `json:"color"`
	Status   Status `json:"status"`
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
func (db *DB) GetProducts(color, status string) ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	list := []Product{}
	builder := Builder()
	selectField := `p.*, c.id "color.id", c.name "color.name", c.code "color.code",
	 s.id "status.id", s.name "status.name", s.code "status.code"`
	where := []string{}

	if len(color) > 0 {
		where = append(where, "c.id = "+color)
	}

	if len(status) > 0 {
		where = append(where, "s.id = "+status)
	}

	query := builder.Select(selectField).
		From("products p").
		Join("colors c ON p.color_id = c.id").
		Join("statuses s ON p.status_id = s.id").
		Where(where, "AND").
		Limit("20").
		ToSql()

	err := db.SelectContext(ctx, &list, query)

	if err != nil {
		return list, nil
	}

	return list, err
}

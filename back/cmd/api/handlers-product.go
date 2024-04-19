package main

import (
	"net/http"

	"github.com/samaelkorn/go-catalog/internal/response"
)

func (app *application) productList(w http.ResponseWriter, r *http.Request) {
	products, err := app.db.GetProducts()

	if err != nil {
		app.serverError(w, r, err)
	}

	data := Response{
		Data: products,
	}

	err = response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

package main

import (
	"net/http"

	"github.com/samaelkorn/go-catalog/internal/response"
)

func (app *application) colorList(w http.ResponseWriter, r *http.Request) {
	list, err := app.db.GetColors()

	if err != nil {
		app.serverError(w, r, err)
	}

	data := Response{
		Data: list,
	}

	err = response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
func (app *application) statusList(w http.ResponseWriter, r *http.Request) {
	list, err := app.db.GetStatuses()

	if err != nil {
		app.serverError(w, r, err)
	}

	data := Response{
		Data: list,
	}

	err = response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

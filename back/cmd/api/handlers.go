package main

import (
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
}

func (app *application) protected(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

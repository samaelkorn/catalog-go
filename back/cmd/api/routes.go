package main

import (
	"net/http"

	"github.com/alexedwards/flow"
)

func (app *application) routes() http.Handler {
	mux := flow.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.Use(app.logAccess)
	mux.Use(app.recoverPanic)
	mux.Use(app.authenticate)

	//ref
	mux.HandleFunc("/colors", app.colorList, "GET")
	mux.HandleFunc("/statuses", app.statusList, "GET")
	//product
	mux.HandleFunc("/products", app.productList, "GET")
	//users
	mux.HandleFunc("/users", app.createUser, "POST")
	mux.HandleFunc("/authentication-tokens", app.createAuthenticationToken, "POST")
	//fileserver
	mux.Handle("/storage/...", http.StripPrefix("/storage/", http.FileServer(http.Dir("./storage"))))

	mux.Group(func(mux *flow.Mux) {
		mux.Use(app.requireAuthenticatedUser)

		mux.HandleFunc("/protected", app.protected, "GET")
	})

	return mux
}

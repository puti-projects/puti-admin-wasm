package main

import (
	"log"
	"net/http"

	"github.com/puti-projects/puti-admin-wasm/internal"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

func main() {
	internal.RunApp()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"/web/css/theme.min.css",
			"/web/icon-set/style.css",
		},
		Scripts: []string{
			"/web/js/jquery.min.js",
			"/web/js/jquery-migrate.min.js",
			"/web/js/bootstrap.bundle.min.js",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

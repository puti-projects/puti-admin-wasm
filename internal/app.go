package internal

import (
	"github.com/puti-projects/puti-admin-wasm/internal/pages/layout"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

func RunApp() {
	app.Route("/", &layout.Body{})

	app.RunWhenOnBrowser()
}

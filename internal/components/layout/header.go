package layout

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	return app.Header().ID("header").Class("navbar navbar-expand-lg navbar-fixed navbar-height navbar-flush navbar-container navbar-bordered").Body(
		app.Div().Class("navbar-nav-wrap").Body(
			app.Div().Class("navbar-brand-wrapper").Body(
				// Logo
				&NavbarLogo{},
			),

			app.Div().Class("navbar-nav-wrap-content-left").Body(
				// Navbar Vertical Toggle
				&NavbarVerticalToggle{},

				// Search Form
			),

			// Secondary Content
			app.Div().Class("navbar-nav-wrap-content-right").Body(
				app.Ul().Class("navbar-nav align-items-center flex-row").Body(
					app.Li().Class("nav-item d-md-none").Text("123"),
				),
			),
		),
	)
}

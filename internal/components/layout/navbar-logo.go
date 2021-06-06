package layout

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type NavbarLogo struct {
	app.Compo
}

func (nl *NavbarLogo) Render() app.UI {
	return app.A().Class("navbar-brand").Href("/").Aria("label", "Front").Body(
		app.Img().Class("navbar-brand-logo").Src("").Alt("Logo"),
		app.Img().Class("navbar-brand-logo-mini").Src("").Alt("Logo"),
	)
}

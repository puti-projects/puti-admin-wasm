package layout

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type NavbarVerticalMenu struct {
	app.Compo

	category string
}

func (nvm *NavbarVerticalMenu) Render() app.UI {
	var showDashboards = ""
	var activeDashboards = ""
	if nvm.category == "dashboards" {
		showDashboards = "show"
		activeDashboards = "active"
	}
	return app.Li().Class("navbar-vertical-aside-has-menu "+showDashboards).Body(
		app.A().Class("js-navbar-vertical-aside-menu-link nav-link nav-link-toggle"+activeDashboards).
			Href("javascript:;").Title("Dashboards").Body(
			app.I().Class("tio-home-vs-1-outlined nav-icon"),
			app.Span().Class("navbar-vertical-aside-mini-mode-hidden-elements text-truncate").Text("Dashboards"),
		),

		app.Ul().Class("js-navbar-vertical-aside-submenu nav nav-sub").Body(
			app.Li().Class("nav-item").Body(
				app.A().Class("nav-link").Href("").Title("Default").Body(
					app.Span().Class("tio-circle nav-indicator-icon"),
					app.Span().Class("text-truncate").Text("Default"),
				),
			),
			app.Li().Class("nav-item").Body(
				app.A().Class("nav-link").Href("").Title("Alternative").Body(
					app.Span().Class("tio-circle nav-indicator-icon"),
					app.Span().Class("text-truncate").Text("Alternative"),
				),
			),
		),
	)
}

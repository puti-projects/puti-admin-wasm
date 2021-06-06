package layout

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type NavbarVerticalToggle struct {
	app.Compo
}

func (nvt *NavbarVerticalToggle) Render() app.UI {
	return app.Button().Type("button").Class("js-navbar-vertical-aside-toggle-invoker close mr-3").Body(
		app.Raw(`
		  <i class="tio-first-page navbar-vertical-aside-toggle-short-align" data-toggle="tooltip" data-placement="right" title="Collapse"></i>
		  <i class="tio-last-page navbar-vertical-aside-toggle-full-align" data-template='<div class="tooltip d-none d-sm-block" role="tooltip"><div class="arrow"></div><div class="tooltip-inner"></div></div>' data-toggle="tooltip" data-placement="right" title="Expand"></i>
		`),
	)
}

package layout

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/puti-projects/puti-admin-wasm/internal/components"
	"github.com/puti-projects/puti-admin-wasm/internal/components/layout"
)

type Body struct {
	app.Compo
}

func (h *Body) Render() app.UI {
	return app.Body().Class("has-navbar-vertical-aside navbar-vertical-aside-show-xl").Body(
		// Search Form
		&components.Search{},

		// HEADER
		&layout.Header{},
		// END HEADER

		// MAIN CONTENT
		// Navbar Vertical
		&layout.Sidebar{},

		//app.Script().Src("/web/js/hs-navbar-vertical-aside.min.js"),
		//app.Script().Src("/web/js/hs-unfold.min.js"),
		//app.Script().Src("/web/js/hs-form-search.min.js"),
		//app.Script().Src("/web/js/hs.core.js"),
		//app.Script().Src("/web/js/main.js"),
	)
}

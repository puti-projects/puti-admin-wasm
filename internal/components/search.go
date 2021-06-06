package components

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Search struct {
	app.Compo
}

func (s *Search) Render() app.UI {
	return app.Div().ID("searchDropdown").Class("hs-unfold-content dropdown-unfold search-fullwidth d-md-none").Body(
		app.Form().Class("input-group input-group-merge input-group-borderless").Body(
			// input prepend
			app.Div().Class("input-group-prepend").Body(
				app.Div().Class("input-group-text").Body(
					app.I().Class("tio-search"),
				),
			),

			// input
			app.Input().Class("form-control rounded-0").Type("search").Placeholder("placeholder").Aria("label", "Search in front"),

			// input append
			app.Div().Class("input-group-append").Body(
				app.Div().Class("input-group-text").Body(
					app.Div().Class("hs-unfold").Body(
						app.Raw(`<a class="js-hs-unfold-invoker" href="javascript:;"
						data-hs-unfold-options='{
						  "target": "#searchDropdown",
						  "type": "css-animation",
						  "animationIn": "fadeIn",
						  "hasOverlay": "rgba(46, 52, 81, 0.1)",
						  "closeBreakpoint": "md"
						}'>`),
						app.I().Class("tio-clear tio-lg"),
					),
				),
			),
		),
	)
}

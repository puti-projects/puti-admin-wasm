package layout

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Sidebar struct {
	app.Compo

	skinDark bool
}

func (s *Sidebar) Render() app.UI {
	var asideClass = "js-navbar-vertical-aside navbar navbar-vertical-aside navbar-vertical navbar-vertical-fixed navbar-expand-xl "
	if !s.skinDark {
		asideClass += "navbar-bordered navbar-light "
	}
	if s.skinDark {
		asideClass += "navbar-dark "
	}

	return app.Aside().Class(asideClass).Body(
		app.Div().Class("navbar-vertical-container").Body(
			app.Div().Class("navbar-vertical-footer-offset").Body(
				app.Div().Class("navbar-brand-wrapper justify-content-between").Body(
					// LOGO
					app.If(s.skinDark,
						&NavbarLogo{},
					).Else(
						&NavbarLogo{},
					),

					// Navbar Vertical Toggle
					app.Button().Type("button").Class("js-navbar-vertical-aside-toggle-invoker navbar-vertical-aside-toggle btn btn-icon btn-xs btn-ghost-dark").Body(
						app.I().Class("tio-clear tio-lg"),
					),
				),

				// Content
				app.Div().Class("navbar-vertical-content").Body(
					app.Ul().Class("navbar-nav navbar-nav-lg nav-tabs").Body(
						&NavbarVerticalMenu{},
					),
				),

				// Footer
				app.Div().Class("navbar-vertical-footer").Body(
					app.Ul().Class("navbar-vertical-footer-list").Body(
						app.Li().Class("navbar-vertical-footer-list-item").Body(
							app.Raw(`<div class="hs-unfold">
										  <a class="js-hs-unfold-invoker btn btn-icon btn-ghost-secondary rounded-circle" href="javascript:;"
											 data-hs-unfold-options='{
											  "target": "#styleSwitcherDropdown",
											  "type": "css-animation",
											  "animationIn": "fadeInRight",
											  "animationOut": "fadeOutRight",
											  "hasOverlay": true,
											  "smartPositionOff": true
											 }'>
											<i class="tio-tune"></i>
										  </a>
										</div>`),
						),

						app.Li().Class("navbar-vertical-footer-list-item").Body(
							app.Raw(`
<div class="hs-unfold">
  <a class="js-hs-unfold-invoker btn btn-icon btn-ghost-secondary rounded-circle" href="javascript:;"
     data-hs-unfold-options='{
      "target": "#otherLinksDropdown",
      "type": "css-animation",
      "animationIn": "slideInDown",
      "hideOnScroll": true
     }'>
    <i class="tio-help-outlined"></i>
  </a>

  <div id="otherLinksDropdown" class="hs-unfold-content dropdown-unfold dropdown-menu navbar-vertical-footer-dropdown">
    <span class="dropdown-header">Help</span>
    <a class="dropdown-item" href="#">
      <i class="tio-book-outlined dropdown-item-icon"></i>
      <span class="text-truncate pr-2" title="Resources &amp; tutorials">Resources &amp; tutorials</span>
    </a>
    <a class="dropdown-item" href="#">
      <i class="tio-command-key dropdown-item-icon"></i>
      <span class="text-truncate pr-2" title="Keyboard shortcuts">Keyboard shortcuts</span>
    </a>
    <a class="dropdown-item" href="#">
      <i class="tio-alt dropdown-item-icon"></i>
      <span class="text-truncate pr-2" title="Connect other apps">Connect other apps</span>
    </a>
    <a class="dropdown-item" href="#">
      <i class="tio-gift dropdown-item-icon"></i>
      <span class="text-truncate pr-2" title="What's new?">What's new?</span>
    </a>
    <div class="dropdown-divider"></div>
    <span class="dropdown-header">Contacts</span>
    <a class="dropdown-item" href="#">
      <i class="tio-chat-outlined dropdown-item-icon"></i>
      <span class="text-truncate pr-2" title="Contact support">Contact support</span>
    </a>
  </div>
</div>`),
						),

						app.Li().Class("navbar-vertical-footer-list-item").Body(
							app.Raw(`<div class="hs-unfold">
  <a class="js-hs-unfold-invoker btn btn-icon btn-ghost-secondary rounded-circle" href="javascript:;"
     data-hs-unfold-options='{
      "target": "#languageDropdown",
      "type": "css-animation",
      "animationIn": "slideInDown",
      "hideOnScroll": true
     }'>
    <img class="avatar avatar-xss avatar-circle" src="/web/flag-icon-css/flags/1x1/us.svg" alt="United States Flag">
  </a>

  <div id="languageDropdown" class="hs-unfold-content dropdown-unfold dropdown-menu navbar-vertical-footer-dropdown">
    <span class="dropdown-header">Select language</span>
    <a class="dropdown-item" href="#">
      <img class="avatar avatar-xss avatar-circle mr-2" src="/web//flag-icon-css/flags/1x1/us.svg" alt="Flag">
      <span class="text-truncate pr-2" title="English">English (US)</span>
    </a>
    <a class="dropdown-item" href="#">
      <img class="avatar avatar-xss avatar-circle mr-2" src="/web/flag-icon-css/flags/1x1/gb.svg" alt="Flag">
      <span class="text-truncate pr-2" title="English">English (UK)</span>
    </a>
    <a class="dropdown-item" href="#">
      <img class="avatar avatar-xss avatar-circle mr-2" src="/web/flag-icon-css/flags/1x1/de.svg" alt="Flag">
      <span class="text-truncate pr-2" title="Deutsch">Deutsch</span>
    </a>
    <a class="dropdown-item" href="#">
      <img class="avatar avatar-xss avatar-circle mr-2" src="/web/flag-icon-css/flags/1x1/dk.svg" alt="Flag">
      <span class="text-truncate pr-2" title="Dansk">Dansk</span>
    </a>
    <a class="dropdown-item" href="#">
      <img class="avatar avatar-xss avatar-circle mr-2" src="/web/flag-icon-css/flags/1x1/it.svg" alt="Flag">
      <span class="text-truncate pr-2" title="Italiano">Italiano</span>
    </a>
    <a class="dropdown-item" href="#">
      <img class="avatar avatar-xss avatar-circle mr-2" src="/web/flag-icon-css/flags/1x1/cn.svg" alt="Flag">
      <span class="text-truncate pr-2" title="中文 (繁體)">中文 (繁體)</span>
    </a>
  </div>
</div>`),
						),
					),
				),
			),
		),
	)
}

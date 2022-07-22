package nav

import (
	"fg_hub/frontend/components/core"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type MenuItem struct {
	app.Compo
	Icon     *core.Icon
	Label    string
	Href     string
	Children []MenuItem
}

func (m *MenuItem) Render() app.UI {
	link := app.A().Href(m.Href).Text(m.Label)
	return app.Li().Body(m.Icon, link)
}

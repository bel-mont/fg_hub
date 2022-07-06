package nav

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type MenuItem struct {
	app.Compo
	Icon  string // TODO: Change to a struct
	Label string
	Href  string
}

func (m *MenuItem) Render() app.UI {
	return app.A().Href(m.Href).Text(m.Label)
}

package nav

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Sidebar struct {
	app.Compo
	items []MenuItem
}

func (s *Sidebar) Render() app.UI {
	return app.Aside().Text("Test")
}

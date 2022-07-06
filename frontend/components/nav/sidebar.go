package nav

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Sidebar struct {
	app.Compo
	Items []MenuItem
}

func (s *Sidebar) Render() app.UI {
	var listItems []app.UI
	for _, item := range s.Items {
		// If I do not call render here, it errors out later for some reason
		listItems = append(listItems, app.Li().Body((&item).Render()))
	}
	return app.Ul().Body(listItems...)
}

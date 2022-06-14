package nav

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Sidebar struct {
	app.Compo
	Items *[]MenuItem
}

func (s *Sidebar) Render() app.UI {
	var listItems []app.UI
	for _, item := range *s.Items {
		listItems = append(listItems, app.Li().Text(item.Label))
	}
	return app.Ul().Body(listItems...)
}

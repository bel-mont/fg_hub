package nav

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Sidebar struct {
	app.Compo
	Items []MenuItem
}

func (s *Sidebar) Render() app.UI {
	listItems := make([]MenuItem, len(s.Items))
	for _, item := range s.Items {
		listItems = append(listItems, item)
	}
	zz := s.Items[0]
	return app.Ul().Body(&zz)
}

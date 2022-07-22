package nav

import (
	"fg_hub/frontend/components/nav"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Menu struct {
	app.Compo
}

func (m *Menu) Render() app.UI {
	sideBar := &nav.Sidebar{Items: m.GetMenuItems()}
	return app.Nav().Body(sideBar)
}

func (m *Menu) GetMenuItems() []nav.MenuItem {
	series := *nav.SeriesMenuItem()
	items := []nav.MenuItem{series}
	return items
}

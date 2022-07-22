package nav

import "fg_hub/frontend/components/core"

func SeriesMenuItem() *MenuItem {
	series := MenuItem{Icon: &core.Icon{Name: "gamepad"}, Label: "Series", Href: "/series"}
	return &series
}

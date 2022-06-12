package nav

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type MenuItem struct {
	app.Compo
	icon  string // Change to a struct
	label string
	href  string
}

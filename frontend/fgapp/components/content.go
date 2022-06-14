package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Content struct {
	app.Compo
}

func (b *Content) Render() app.UI {
	return app.Section().Text("Application Content goes here")
}

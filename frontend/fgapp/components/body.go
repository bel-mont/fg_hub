package components

import (
	"fg_hub/frontend/fgapp/components/nav"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Body struct {
	app.Compo
}

func (b *Body) Render() app.UI {
	menu := &nav.Menu{}
	content := &Content{}
	return app.Main().Body(menu, content)
}

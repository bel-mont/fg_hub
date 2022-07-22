package core

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Icon struct {
	app.Compo
	Name string
}

func (i *Icon) Render() app.UI {
	return app.I().Class(fmt.Sprintf("fa-solid fa-%s", i.Name))
}

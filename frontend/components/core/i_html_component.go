package core

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type IHtmlComponent interface {
	Render() app.UI
}

package core

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type IHtml interface {
	Render() app.UI
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

func (f *hello) OnPreRender(ctx app.Context) {
	fmt.Println("component prerendered, server side")
}

func (f *hello) OnMount(ctx app.Context) {
	fmt.Println("component mounted into dom")
}

func (f *hello) OnNav(ctx app.Context) {
	fmt.Println("component navigated client side")
}

func (f *hello) OnDismount() {
	fmt.Println("component dismounted from dom, client side")
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.Div().ID("app").Body(
		app.Aside().Class("side-menu").Body(
			app.Div().Class("side-menu-title").Body(
				app.H1().Body(
					app.Text("FG Hub"),
				),
			),
			app.Div().Class("side-menu-game-select").Body(
				app.H2().Body(
					app.Text("Select a Game"),
				),
			),
			app.Div().Class("side-menu-contextual-menu").Body(
				app.P().Body(
					app.Text("Game Dependent Contextual Menu"),
				),
			),
		),
		app.Main().Class("app-body").Body(
			app.H1().Body(
				app.Text("Welcome to FG Hub"),
			),
			app.P().Body(
				app.Text("To start, pick a game from the menu on the left"),
			),
		),
	)
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &hello{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Hello kind sir",
		Description: "A Hello World! example",
		Title:       "FG Hub",
		Styles: []string{
			"/web/css/normalize.css",
			"/web/css/main.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

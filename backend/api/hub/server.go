package main

import (
	"fg_hub/backend/api/hub/graph"
	"fg_hub/backend/api/hub/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/bel-mont/fg_hub/backend/db/rdb"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "5432"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	rdb.InitDB()
	rdb.Migrate()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

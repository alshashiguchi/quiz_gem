package main

import (
	"alshashiguchi/quiz_gem/graph"
	"alshashiguchi/quiz_gem/graph/generated"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"

	configurations "alshashiguchi/quiz_gem/core"
	database "alshashiguchi/quiz_gem/db/mysql"
	auth "alshashiguchi/quiz_gem/middleware/auth"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	config := configurations.New()

	port := config.PortServer.Port
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB(config)
	database.Migrate(config)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

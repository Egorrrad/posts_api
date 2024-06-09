package main

import (
	"GraphQL_api/graph"
	"GraphQL_api/graph/model"
	"GraphQL_api/internal"
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"
const defaultStorage = "in-memory"

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = defaultPort
	}
	storage, exists := os.LookupEnv("STORAGE")
	if !exists {
		storage = defaultStorage
	}

	store := internal.NewDataStorage(storage)

	resolver := &graph.Resolver{
		Store: store,
	}

	config := graph.Config{Resolvers: resolver}

	config.Directives.MaxLengthComment = func(ctx context.Context, obj interface{}, next graphql.Resolver, max int) (res interface{}, err error) {
		copmment, ok := obj.(*model.Comment)
		text := copmment.Text
		if !ok {
			return nil, errors.New("не удается получить текст")
		}

		if len(text) > max {
			return nil, errors.New("Текст превышает максимальную длину")
		}

		return next(ctx)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//  go run ./server/server.go

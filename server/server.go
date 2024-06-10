package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"log"
	"net/http"
	"os"
	"posts_api/graph"
	"posts_api/graph/model"
	"posts_api/internal"
	"posts_api/internal/postgres"
	"time"

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

	store, db := internal.NewDataStorage(storage)

	if storage == "postgres" {
		fmt.Println("Успешный запуск базы данных postgres")
		defer postgres.CloseDB(db)
	}

	resolver := &graph.Resolver{
		Store:       store,
		Comments:    make(map[int][]chan *model.Comment),
		NewComments: make(chan *model.Comment, 100),
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

	go resolver.HandleNewComments()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//  go run ./server/server.go

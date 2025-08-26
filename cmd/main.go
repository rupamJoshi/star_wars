package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rupam_joshi/star_wars/config"
	"github.com/rupam_joshi/star_wars/graph"
	"github.com/rupam_joshi/star_wars/repo"
	"github.com/rupam_joshi/star_wars/service"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	config := config.NewConfig("./config/config.yaml")
	if config == nil {
		log.Fatal("config not found")
	}
	// fmt.Printf("%+v", config)
	dbConfig := config.DBConfig
	dbHost := dbConfig.Host
	dbPort := dbConfig.Port
	dbURL := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	repo, err := repo.New(dbURL, dbConfig.DBName, dbConfig.CollectionName)
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}
	service := service.NewStarWarsService(*config, repo)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Service: service,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	host := config.Host
	port := config.Port
	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}

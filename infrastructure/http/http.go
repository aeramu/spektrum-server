package main

import (
	"context"
	"log"
	"net/http"

	resolver "github.com/aeramu/spektrum-server/implementation/graphql.resolver"
	repository "github.com/aeramu/spektrum-server/implementation/mongodb.repository"
	"github.com/aeramu/spektrum-server/interactor"
	"github.com/friendsofgo/graphiql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	context := context.WithValue(context.Background(), "request", map[string]string{
		"token": "18118041",
	})

	schema := graphql.MustParseSchema(resolver.Schema, resolver.Constructor{
		Context: context,
		Interactor: interactor.Constructor{
			Repository: repository.New(),
		}.New(),
	}.New())
	http.Handle("/", corsMiddleware(&relay.Handler{Schema: schema}))

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/")
	if err != nil {
		panic(err)
	}
	http.Handle("/graphiql", graphiqlHandler)

	log.Println("Server ready at 8000")
	log.Println("Graphiql ready at 8000/graphiql")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

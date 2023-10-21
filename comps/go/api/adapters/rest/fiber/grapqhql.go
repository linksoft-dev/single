package fiber

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

var schema *graphql.Schema

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootQuery",
	Fields: graphql.Fields{},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootMutation",
	Fields: graphql.Fields{},
})

// createGraphQlSchema criato o schema vindo das apps
func createGraphQlSchema() {
	if len(rootQuery.Fields()) == 0 && len(rootMutation.Fields()) == 0 {
		return
	}

	schemaConfig := graphql.SchemaConfig{}
	if len(rootQuery.Fields()) > 0 {
		schemaConfig.Query = rootQuery
	}

	if len(rootMutation.Fields()) > 0 {
		schemaConfig.Mutation = rootMutation
	}

	var err error
	schema2, err := graphql.NewSchema(schemaConfig)
	schema = &schema2

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
}

func handler1(w http.ResponseWriter, req *http.Request) {
	h := handler.New(&handler.Config{
		Schema:     schema,
		Pretty:     true,
		Playground: true,
	})
	h.ServeHTTP(w, req)
	fmt.Fprint(w, "something")
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

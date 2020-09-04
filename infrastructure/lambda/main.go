package main

import (
	"context"
	"encoding/json"

	resolver "github.com/aeramu/spektrum-server/implementation/graphql.resolver"
	repository "github.com/aeramu/spektrum-server/implementation/mongodb.repository"
	"github.com/aeramu/spektrum-server/interactor"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graph-gophers/graphql-go"
)

var i interactor.Interactor = nil

func main() {
	lambda.Start(handler)
}

func initInteractor() interactor.Interactor {
	if i != nil {
		return i
	}
	i = interactor.Constructor{
		Repository: repository.New(),
	}.New()
	return i
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//convert request body to json
	var parameter struct {
		Query         string
		OperationName string
		Variables     map[string]interface{}
	}
	json.Unmarshal([]byte(request.Body), &parameter)

	//add token from header
	context := context.WithValue(ctx, "request", request.Headers)

	//graphql execution
	schema := graphql.MustParseSchema(resolver.Schema, resolver.Constructor{
		Context:    context,
		Interactor: initInteractor(),
	}.New())
	response := schema.Exec(context, parameter.Query, parameter.OperationName, parameter.Variables)
	responseJSON, _ := json.Marshal(response)

	//response
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil
}

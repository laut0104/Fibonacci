package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func fib(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	a := big.NewInt(0)
	b := big.NewInt(1)

	n, err := strconv.Atoi(request.QueryStringParameters["n"])
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	for i := 0; i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}

	log.Println(a, b)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       `{"result":` + fmt.Sprintf("%s", a) + `}`,
	}, nil
}

func main() {
	lambda.Start(fib)
}

package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	BadRequest = events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       `{"status": 400, "message": "` + http.StatusText(http.StatusBadRequest) + `."}`,
	}
	MethodNotAllowed = events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       `{"status": 405, "message": "` + http.StatusText(http.StatusMethodNotAllowed) + `."}`,
	}
)

func fib(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(request)
	switch request.HTTPMethod {
	case "GET":
		a := big.NewInt(0)
		b := big.NewInt(1)

		q := request.QueryStringParameters["n"]
		if q == "" {
			return BadRequest, nil
		}

		n, err := strconv.Atoi(q)
		if err != nil {
			log.Println(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"status": 400, "message":` + err.Error() + `}`,
			}, err
		}
		if n < 1 {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"status": 400, "message": "n is a natural number."}`,
			}, nil
		}

		for i := 0; i < n; i++ {
			a.Add(a, b)
			a, b = b, a
		}

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       `{"result":` + fmt.Sprintf("%s", a) + `}`,
		}, nil
	default:
		return MethodNotAllowed, nil
	}
}

func main() {
	lambda.Start(fib)
}

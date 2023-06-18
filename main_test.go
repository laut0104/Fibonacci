package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// GET /fib
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "GET",
			},
			expect: `{"status": 400, "message": "Bad Request."}`,
			err:    nil,
		},
		{
			//GET /fib?n=1
			request: events.APIGatewayProxyRequest{
				HTTPMethod:            "GET",
				QueryStringParameters: map[string]string{"n": "1"},
			},
			expect: `{"result":1}`,
			err:    nil,
		},
		{
			//GET /fib?n=99
			request: events.APIGatewayProxyRequest{
				HTTPMethod:            "GET",
				QueryStringParameters: map[string]string{"n": "99"},
			},
			expect: `{"result":218922995834555169026}`,
			err:    nil,
		},
		{
			//GET /fib?n=0
			request: events.APIGatewayProxyRequest{
				HTTPMethod:            "GET",
				QueryStringParameters: map[string]string{"n": "0"},
			},
			expect: `{"status": 400, "message": "n is a natural number."}`,
			err:    nil,
		},
		{
			//GET /fib?n=-1
			request: events.APIGatewayProxyRequest{
				HTTPMethod:            "GET",
				QueryStringParameters: map[string]string{"n": "-1"},
			},
			expect: `{"status": 400, "message": "n is a natural number."}`,
			err:    nil,
		},
		{
			//POST /fib?n=1
			request: events.APIGatewayProxyRequest{
				HTTPMethod:            "POST",
				QueryStringParameters: map[string]string{"n": "1"},
			},
			expect: `{"status": 405, "message": "Method Not Allowed."}`,
			err:    nil,
		},
	}

	for _, test := range tests {
		response, err := fib(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}

}

package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"number": "5",
		},
	}

	expectedResponse := Response{
		Result: 10,
	}

	ctx := context.Background()
	response, err := handler(ctx, request)

	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	var actualResponse Response
	err = json.Unmarshal([]byte(response.Body), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

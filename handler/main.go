package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/matishsiao/goInfo"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	gi, err := goInfo.GetInfo()

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf(err.Error()),
			StatusCode: 200,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf(gi.String() + "V2"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

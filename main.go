package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"strconv"
)

type TriggerEvent struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

func HandleRequest(ctx lambdacontext.LambdaContext, request TriggerEvent) (string, error) {
	data, err := GetPagePdf(request.Url)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	success, saveErr := SaveFile(data, request.Key)
	if saveErr != nil {
		fmt.Println(saveErr.Error())
		return "", saveErr
	}
	return strconv.FormatBool(success), nil
}

func main() {
	lambda.Start(HandleRequest)
}

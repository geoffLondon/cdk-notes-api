package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/strategy/ctxmissing"
	"github.com/aws/aws-xray-sdk-go/xray"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
}

func main() {
	if err := xray.Configure(xray.Config{
		ContextMissingStrategy: ctxmissing.NewDefaultLogErrorStrategy(),
	}); err != nil {
		panic(err)
	}

	container, err := InitializeAppsyncContainer()
	if err != nil {
		panic(err)
	}

	lambda.Start(container.Resolver().Handle)
}

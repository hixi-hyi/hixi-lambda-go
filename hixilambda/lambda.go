package hixilambda

import (
	"flag"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/aws-client-go/awsclient"
	"github.com/hixi-hyi/localstack-go/localstack"
	"github.com/hixi-hyi/logone-lambda-go/lambdalog"
)

var LocalStackDomain = "localstack"

type Lambda struct {
	AwsSession *session.Session
	LogManager *lambdalog.Manager
}

func (l *Lambda) Init() {
	l.AwsSession = session.Must(session.NewSession())
	l.LogManager = lambdalog.NewManagerDefault()

	if isSamLocal() || isTest() {
		ls := localstack.New(&localstack.Config{Domain: LocalStackDomain})
		awsclient.UseLocalStack(ls)
		l.LogManager.Config.JsonIndent = true
	}
}

func isSamLocal() bool {
	return os.Getenv("AWS_SAM_LOCAL") == "true"
}

func isTest() bool {
	return flag.Lookup("test.v") != nil
}

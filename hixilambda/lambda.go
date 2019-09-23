package hixilambda

import (
	"flag"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/aws-client-go/awsclient"
	"github.com/hixi-hyi/localstack-go/localstack"
	"github.com/hixi-hyi/logone-lambda-go/lambdalog"
)

var LocalStackDomain = ""

type Lambda struct {
	AwsSession   *session.Session
	LogManager   *lambdalog.Manager
	Environments Environments
}

func (l *Lambda) Init() {
	l.AwsSession = session.Must(session.NewSession())
	l.LogManager = lambdalog.NewManagerDefault()
	l.Environments = Environments{}

	if isSamLocal() || isTest() || hasLocalStackDomain() {
		// In generally, I should not put the code for testing in here.
		// But, I should put it if I want to use `sam local`.
		var domain string
		if hasLocalStackDomain() {
			domain = LocalStackDomain
		} else if isSamLocal() {
			domain = "localstack"
		} else if isTest() {
			domain = "localhost"
		} else {
			panic("sould be set hixilambda.LocalStackDomain")
		}
		ls := localstack.New(&localstack.Config{Domain: domain})
		awsclient.UseLocalStack(ls)
		l.LogManager.Config.JsonIndent = true
	}
}

func hasLocalStackDomain() bool {
	return LocalStackDomain != ""
}

func isSamLocal() bool {
	return os.Getenv("AWS_SAM_LOCAL") == "true"
}

func isTest() bool {
	return flag.Lookup("test.v") != nil
}

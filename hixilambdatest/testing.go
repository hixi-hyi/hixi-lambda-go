package hixilambdatest

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/aws-client-go/awsclient"
	"github.com/hixi-hyi/hixi-lambda-go/hixilambda"
	"github.com/hixi-hyi/localstack-go/localstack"
	"github.com/hixi-hyi/logone-lambda-go/lambdalog"
)

var LocalStackDomain = ""

func RunWithContext(t *testing.T, name string, f func(t *testing.T, ctx context.Context)) {
	ctx := lambdacontext.NewContext(context.Background(), &lambdacontext.LambdaContext{})
	{
		var domain string
		if LocalStackDomain != "" {
			domain = LocalStackDomain
		} else {
			domain = "localhost"
		}
		ls := localstack.New(&localstack.Config{Domain: domain})
		awsclient.UseLocalStack(ls)
	}
	ctx = hixilambda.NewContextWithAwsSession(ctx, session.Must(session.NewSession()))

	lm := lambdalog.NewManagerDefault()
	lm.Config.JsonIndent = true
	log, d := lm.Recording(ctx)
	defer d()
	ctx = lambdalog.NewContextWithLogger(ctx, log)
	t.Run(name, func(t *testing.T) {
		f(t, ctx)
	})
}


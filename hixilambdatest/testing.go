package hixilambdatest

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/hixi-hyi/logone-lambda-go/lambdalog"
)

func RunWithContext(t *testing.T, name string, f func(t *testing.T, ctx context.Context)) {
	ctx := lambdacontext.NewContext(context.Background(), &lambdacontext.LambdaContext{})
	lm := lambdalog.NewManagerDefault()
	lm.Config.JsonIndent = true
	log, d := lm.Recording(ctx)
	defer d()
	ctx = lambdalog.NewContextWithLogger(ctx, log)
	t.Run(name, func(t *testing.T) {
		f(t, ctx)
	})
}


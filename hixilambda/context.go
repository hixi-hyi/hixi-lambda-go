package hixilambda

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
)

type key struct{}

var sessionKey = &key{}

func NewContextWithAwsSession(parent context.Context, sess *session.Session) context.Context {
	return context.WithValue(parent, sessionKey, sess)
}

func AwsSessionFromContext(ctx context.Context) (*session.Session, bool) {
	sess, ok := ctx.Value(sessionKey).(*session.Session)
	return sess, ok
}


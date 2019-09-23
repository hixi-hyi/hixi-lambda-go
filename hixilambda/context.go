package hixilambda

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
)

type key int

const (
	sessionKey key = iota
	environmentKey
)

func NewContextWithAwsSession(parent context.Context, sess *session.Session) context.Context {
	return context.WithValue(parent, sessionKey, sess)
}

func AwsSessionFromContext(ctx context.Context) (*session.Session, bool) {
	sess, ok := ctx.Value(sessionKey).(*session.Session)
	return sess, ok
}

type Environments map[string]interface{}

func NewContextWithEnvironments(parent context.Context, envs Environments) context.Context {
	return context.WithValue(parent, environmentKey, envs)
}

func EnvironmentsFromContext(ctx context.Context) (Environments, bool) {
	envs, ok := ctx.Value(environmentKey).(Environments)
	return envs, ok
}

func (e Environments) MustGetString(key string) string {
	return e[key].(string)
}

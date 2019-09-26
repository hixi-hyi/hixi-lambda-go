package hixilambda

import (
	"context"
	"fmt"
	"os"
	"strings"

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

func NewEnvironments() Environments {
	return Environments{}
}

func NewContextWithEnvironments(parent context.Context, envs Environments) context.Context {
	return context.WithValue(parent, environmentKey, envs)
}

func EnvironmentsFromContext(ctx context.Context) (Environments, bool) {
	envs, ok := ctx.Value(environmentKey).(Environments)
	return envs, ok
}

func (e Environments) MustGetString(key string) string {
	got, ok := e[key]
	if !ok {
		panic(fmt.Sprintf("key:%s is not set in environments", key))
	}
	ret, ok := got.(string)
	if !ok {
		panic(fmt.Sprintf("key:%s is not string value: %v", key, got))
	}
	return ret
}

func (e Environments) LoadEnvOnlyPrefix(prefix string) {
	osenvs := os.Environ()
	for _, osenv := range osenvs {
		splits := strings.SplitN(osenv, "=", 2)
		key := splits[0]
		value := splits[1]
		if strings.HasPrefix(key, prefix) {
			e[key] = value
		}
	}
}
func (e Environments) MustLoadEnv(key string) {
	got, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("key:%s does not exist in the environments", key))
	}
	e[key] = got
}

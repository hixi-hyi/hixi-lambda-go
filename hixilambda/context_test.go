package hixilambda_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/hixi-lambda-go/hixilambda"
	"github.com/stretchr/testify/assert"
)

func TestAwsSession(t *testing.T) {

	ctx := context.Background()
	sess := session.Must(session.NewSession())

	t.Run("not set", func(t *testing.T) {
		_, ok := hixilambda.AwsSessionFromContext(ctx)
		assert.Equal(t, ok, false)
	})

	ctx = hixilambda.NewContextWithAwsSession(ctx, sess)
	t.Run("set before", func(t *testing.T) {
		got, ok := hixilambda.AwsSessionFromContext(ctx)
		assert.Equal(t, ok, true)
		assert.Equal(t, got, sess)
	})
}

func TestEnvironments(t *testing.T) {

	ctx := context.Background()
	envs := hixilambda.Environments{}

	t.Run("not set", func(t *testing.T) {
		_, ok := hixilambda.EnvironmentsFromContext(ctx)
		assert.Equal(t, ok, false)
	})

	ctx = hixilambda.NewContextWithEnvironments(ctx, envs)
	t.Run("set before", func(t *testing.T) {
		got, ok := hixilambda.EnvironmentsFromContext(ctx)
		assert.Equal(t, ok, true)
		assert.Equal(t, got, envs)
	})
}

func TestEnvironmentsFunction(t *testing.T) {
	t.Run("MustGetString", func(t *testing.T) {
		envs := hixilambda.Environments{
			"String": "string",
			"Int":    1,
		}
		assert.Equal(t, envs.MustGetString("String"), "string")
		assert.PanicsWithValue(t, fmt.Sprintf("key:Int is not string value: %v", 1), func() { envs.MustGetString("Int") })
		assert.PanicsWithValue(t, "key:NoValue is not set in environments", func() { envs.MustGetString("NoValue") })
	})
	t.Run("LoadEnvOnlyPrefix", func(t *testing.T) {
		envs := hixilambda.Environments{}
		os.Setenv("HIXI_VALUE", "hixi")
		os.Setenv("NOTHIXI_VALUE", "hixi")
		envs.LoadEnvOnlyPrefix("HIXI")
		assert.Equal(t, envs.MustGetString("HIXI_VALUE"), "hixi")
		assert.Panics(t, func() { envs.MustGetString("NOHIXI_VALUE") })
	})
	t.Run("MustLoadEnv", func(t *testing.T) {
		envs := hixilambda.Environments{}
		os.Setenv("HIXI_VALUE", "hixi")
		envs.MustLoadEnv("HIXI_VALUE")
		assert.Equal(t, envs.MustGetString("HIXI_VALUE"), "hixi")

		assert.PanicsWithValue(t, "key:HIXI_NOT_EXISTS_VALUE does not exist in the environments", func() { envs.MustLoadEnv("HIXI_NOT_EXISTS_VALUE") })
	})
}

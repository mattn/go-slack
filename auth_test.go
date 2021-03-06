package slack_test

import (
	"context"
	"testing"

	"github.com/lestrrat/go-slack"
	"github.com/stretchr/testify/assert"
)

func TestAuthTest(t *testing.T) {
	if !hasTestSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)
	res, err := c.Auth().Test(ctx)
	if !assert.NoError(t, err, "Auth.Test failed") {
		return
	}
	t.Logf("%#v", res)
}

func TestAuthRevoke(t *testing.T) {
	if !hasTestSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)
	err := c.Auth().Revoke(ctx, true)
	if !assert.NoError(t, err, "Auth.Revoke failed") {
		return
	}
}


package resolver_test

import (
	"context"
	"github.com/spewerspew/spew"
	"github.com/stretchr/testify/require"
	"testing"
	"util/resolver"
)

func TestGetAddress(t *testing.T) {
	ctx := context.Background()
	addrs, err := resolver.GetServiceAddr(ctx, "consul")
	require.NoError(t, err)
	spew.Dump(addrs)
}

func TestGetRandomServiceAddr(t *testing.T) {
	ctx := context.Background()
	addrs, err := resolver.GetRandomAddr(ctx, "consul")
	require.NoError(t, err)
	spew.Dump(addrs)
}

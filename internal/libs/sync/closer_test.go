package sync_test

import (
	"testing"
	"time"

	tmsync "github.com/number571/tendermint/internal/libs/sync"
	"github.com/stretchr/testify/require"
)

func TestCloser(t *testing.T) {
	closer := tmsync.NewCloser()

	var timeout bool

	select {
	case <-closer.Done():
	case <-time.After(time.Second):
		timeout = true
	}

	for i := 0; i < 10; i++ {
		closer.Close()
	}

	require.True(t, timeout)
	<-closer.Done()
}

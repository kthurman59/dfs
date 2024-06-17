package p2p

import (
	"testing"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr)

	// Server
	// tr.Start()
}
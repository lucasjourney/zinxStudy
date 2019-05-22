package net

import "testing"

func TestServer(T *testing.T) {
	s := NewServer("zinx v0.1")

	s.Serve()
}

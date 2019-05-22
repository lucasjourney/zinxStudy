package main

import (
	"szinx/net"
)

func main() {
	s := net.NewServer("zinx-v0.1")

	s.Serve()
}

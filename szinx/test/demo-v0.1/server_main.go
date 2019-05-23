package demo_v0_1

import (
	"zinxStudy/szinx/net"
)

func main() {
	s := net.NewServer("zinx-v0.1")

	s.Serve()
}

package main

import (
	"github.com/adamb/goflow/backends"
	"github.com/adamb/goflow/frontends"
	"net"
)

func main() {
	nf := frontends.Netflow{
		BindAddr: net.ParseIP("127.0.0.1"),
		BindPort: 9999,
	}
	b := backends.Dump{}
	nf.Start(b)
}

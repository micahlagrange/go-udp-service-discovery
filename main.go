package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/micahlagrange/go-udp-service-discovery/server"
)

func main() {

	// flags
	portArg := flag.String("port", "65535", "The port to bind to")
	flag.Parse()

	fmt.Printf("Starting udp server on port %s\n", *portArg)

	port, err := strconv.Atoi(*portArg)
	if err != nil {
		panic(err)
	}

	server.Listen(port)
}

package server

import (
	"fmt"
	"net"
)

// Listen - Create a udp server listening on port
func Listen(port int) {
	ServerAddress, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}

	ServerConnection, err := net.ListenUDP("udp", ServerAddress)
	if err != nil {
		panic(err)
	}
	defer ServerConnection.Close()
	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConnection.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
		ServerConnection.WriteToUDP([]byte("ack"), addr)
	}

}

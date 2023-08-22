package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
	 	fmt.Println("Failed to bind to port 6379")
	 	os.Exit(1)
	}

	conn, err := l.Accept()

	if err != nil {
	 	fmt.Println("Error accepting connection: ", err.Error())
	 	os.Exit(1)
	}

	fmt.Printf("Listening on host: %s, port: %s\n", "0.0.0.0", "6379")

	buf := make([]byte, 1024);

	for i := 0; i < 10; i++ {
		size, readerr  := conn.Read(buf);
		fmt.Printf("Read : %d\n", size);
		if readerr != nil {
			fmt.Println("Error reading: ", readerr.Error())
			break;
		}

		conn.Write([]byte("+PONG\r\n"));

	}


}

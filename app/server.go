package main

import (
	"fmt"
	"net"
	"os"
)


func handle(conn net.Conn) {

	fmt.Printf("Listening on host: %s, port: %s\n", "0.0.0.0", "6379")

	buf := make([]byte, 1024);

	for true {

		size, readerr  := conn.Read(buf);

		fmt.Printf("Read : %d\n", size);
		if readerr != nil {
			fmt.Println("Error reading: ", readerr.Error())
			break;
		}

		conn.Write([]byte("+PONG\r\n"));

	}
}


func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")


	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
	 	fmt.Println("Failed to bind to port 6379")
	 	os.Exit(1)
	}


	for true {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handle(conn);
	}



}

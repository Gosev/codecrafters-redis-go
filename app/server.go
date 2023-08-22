package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)


func handle(conn net.Conn) {

	fmt.Printf("Listening on host: %s, port: %s\n", "0.0.0.0", "6379")

	buf := make([]byte, 1024);

	for true {

		_, readerr  := conn.Read(buf);

		//fmt.Printf("Reading input : %s\n", buf);

		if readerr != nil {
			fmt.Println("Error reading: ", readerr.Error())
			break;
		}

		str := string(buf);


		array_size := buf[1] - '0';
		fmt.Printf("Message word count : %d\n", array_size)

		stringBits := strings.Split(str, "\r\n")



		firstSignificantItem := strings.ToLower(stringBits[2]);

		switch firstSignificantItem {

			case "ping":
				conn.Write([]byte("+PONG\r\n"));

			case "echo":

				msg := "+" + stringBits[4] + "\r\n";
				conn.Write([]byte(msg));

			default:
				conn.Write([]byte("+OK\r\n"));

		}

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

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)


func reply(conn net.Conn, msg string) {
	ret := "+" + msg + "\r\n";
	conn.Write([]byte(ret));
}


func handle(conn net.Conn, memoryMap map[string]string) {

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
				reply(conn, "PONG");

			case "set":
				key :=  stringBits[4]
				value :=  stringBits[6]
				memoryMap[key] = value
				reply(conn, "OK");

			case "get":
				key :=  stringBits[4]
				reply(conn, memoryMap[key]);


			case "echo":
				reply(conn, stringBits[4]);

			default:
				reply(conn, "OK");


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

	memoryMap := make(map[string]string)


	for true {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handle(conn, memoryMap);
	}



}

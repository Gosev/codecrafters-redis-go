package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type val struct {
	value string
	createdAt  int
	expiry int64
}

func reply(conn net.Conn, msg string) {

	if (msg == "") {
		conn.Write([]byte("_\r\n"));
		return;
	}

	conn.Write([]byte("+" + msg + "\r\n"));
}


func newValue(value string, expiry int) *val {

		p := val{value: value}

		if (expiry > 0) {
			p.expiry = time.Now().UnixMilli() + int64(expiry);
		} else {
			p.expiry = 0;
		}
		return &p
}

func getValue(key string, memoryMap map[string]*val) string {

	val := memoryMap[key];

	if (val == nil) {
		return ""
	}

	if (val.expiry > 0 && val.expiry < time.Now().UnixMilli()) {
		memoryMap[key] = nil;
		return "";
	}

	return val.value;
}

func handle(conn net.Conn, memoryMap map[string]*val) {

	fmt.Printf("Listening on host: %s, port: %s\n", "0.0.0.0", "6379")

	buf := make([]byte, 1024);

	for true {

		_, readerr  := conn.Read(buf);

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
        var expiry = 0;

				if (len(stringBits) >= 10) {

					subCommand := strings.ToLower(stringBits[8]);

					// px is ttl in milliseconds
					if (subCommand  == "px") {
						exp, convError :=  strconv.Atoi(stringBits[10])

						if convError != nil {
							fmt.Println("Error reading expiry: ", convError.Error())
							exp = 0
						}
						expiry = exp;
					}
					fmt.Printf("Expiry : %d\n", expiry)
				}

				memoryMap[key] = newValue(value, expiry)
				reply(conn, "OK");

			case "get":
				key :=  stringBits[4]
				reply(conn, getValue(key, memoryMap));


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

	memoryMap := make(map[string]*val)

	for true {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handle(conn, memoryMap);
	}



}

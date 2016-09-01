package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	for {
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text + "\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}
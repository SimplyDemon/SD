package main

import (
	"fmt"
	"net"
	"os"
	"bufio"

	"strings"
)

const (
	tcpProtocol = "tcp4"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var listenAddr = &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}

func listen() {
	l, err := net.ListenTCP(tcpProtocol, listenAddr); checkErr(err)
	fmt.Println("Listen port: ", l.Addr().(*net.TCPAddr).Port)

	fmt.Println("we are ready, GJ!")
	c, err := l.AcceptTCP(); checkErr(err)

	for {

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println("Message from:", c.RemoteAddr())
		fmt.Println("message is ", message)


		if strings.Trim(message, "\n"+"\r") == "HELLO" {
			c.Write([]byte("WORLD" + "\n"))
		} else {

			c.Write([]byte("BAD REQUEST" + "\n"))
		}
	}

}

func main() {
	listen()
}
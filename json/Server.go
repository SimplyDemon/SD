package main

import (
	"fmt"
	"net"
	"bufio"
)


var listenAddr = &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}

const (
	tcpProtocol = "tcp4"
)

func main() {

	listen()

}

func listen() {

	l, err := net.ListenTCP(tcpProtocol, listenAddr); checkErr(err)
	fmt.Println("we are ready, GJ!")
	c, err := l.AcceptTCP(); checkErr(err)

	for {

		message, _ := bufio.NewReader(c).ReadString('\n')
		decoded := decode(message)
		LogText("Loop started")

		if isJson(message) && isValidIdentifier(decoded.Identifier) && isValidPlatform(decoded.Platform) && isAnyText(decoded.Text) {

			writeLog(decoded.Identifier, decoded.Platform, decoded.Text)

			c.Write([]byte("Added to log file" + "\n"))

		} else {

			c.Write([]byte("Please, enter your identifier, platform, text correctly! " + "\n"))
			LogText("we got a problems!")
		}
		LogText("Loop finished" + "\n" + "\n")
	}
}
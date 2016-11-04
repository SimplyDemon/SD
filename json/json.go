package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"os"
	"net"
	"bufio"
	"strings"
	"log"
	"strconv"
	"time"
)

type Request struct {
	Identifier string     `json:"identifier"`
	Platform   string       `json:"platform"`
	Text       string           `json:"text"`
}

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

func writeLog(identifier string, platform string, text string) {

	f, err := os.OpenFile("./json/" + time.Now().Format("2006.01.02") + ".logs", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Print("\n" + "Identifier is: " + identifier + ". Platform  is: " + platform + ". Text is: " + text + "." + "\n")
}

func LogText(funcName string) {
	f, err := os.OpenFile("./debugFile.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(funcName)
}

func isAnyText(text string) bool {
	LogText("DEBUG: func isAnyText with arguments " + text + " started!")
	LogText("DEBUG: func isAnyText is : " + strconv.FormatBool(len(text) > 0))
	return len(text) > 0
}
func isValidPlatform(platform string) bool {
	LogText("DEBUG: func isValidPlatform with arguments " + platform + " started!")
	switch platform {
	case
		"browser", "ios", "android", "cli", "ms":
		LogText("DEBUG: func isValidPlatform is: true")
		return true
	}

	LogText("DEBUG: func isValidPlatform is : false")
	return false
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isValidIdentifier(ID string) bool {
	LogText("DEBUG: func isValidIdentifier with arguments " + ID + " started!")
	matched, err := regexp.MatchString("^[A-Fa-f\\d]{5}$", ID)
	LogText("DEBUG: func isValidIdentifier is : " + strconv.FormatBool(matched))
	return matched || err != nil
}
func encode(identifier string, platform string, text string) string {

	req := &Request{
		Identifier:   identifier,
		Platform:   platform,
		Text:   text}

	convert, _ := json.Marshal(req)
	return string(convert)
}

func decode(jsonData string) Request {

	res := Request{}
	unmarsh := []byte(jsonData)
	json.Unmarshal(unmarsh, &res)
	return res

}

func isJson(s string) bool {
	LogText("DEBUG: func isJson with arguments " + s + " started!")
	trim := strings.TrimSpace(s)
	check := string(trim[0]) == "{" && string(trim[len(trim) - 1]) == "}"

	LogText("DEBUG: func isJson is : " + strconv.FormatBool(check))
	return check
}
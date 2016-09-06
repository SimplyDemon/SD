package main

import (
	"encoding/json"
	"fmt"
)

var identifier, platform, text = "aaccb", "ios", "cedax"

type Request struct {
	Identifier string     `json:"identifier"`
	Platform   string       `json:"platform"`
	Text       string           `json:"text"`
}

func main() {

	encode(identifier, platform, text)
	theString := encode(identifier, platform, text)
	decode(theString)

}

func encode(identifier string, platform string, text string) []byte {

	curRequest := &Request{
		Identifier:   identifier,
		Platform:   platform,
		Text:   text}

	jsonData, _ := json.Marshal(curRequest)
	fmt.Println(string(jsonData))
	return jsonData
}

func decode(jsonData []byte) {

	res := Request{}
	json.Unmarshal(jsonData, &res)
	fmt.Println("decode", res)

}

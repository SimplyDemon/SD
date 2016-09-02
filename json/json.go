package main

import (
	"encoding/json"
	"fmt"
)






var identifier, platform, text ="aaccb", "ios", "cedax"
type our struct {
	Identifier string     `json:"identifier"`
	Platform   string       `json:"platform"`
	Text       string           `json:"text"`
}

func main () {


	encode(identifier, platform, text)
	theString := encode(identifier, platform, text)
	decode(theString)

}

func encode(identifier string, platform string, text string) []byte {


	our2 := &our{
		Identifier:   identifier,
		Platform:   platform,
		Text:   text}

	convert, _ := json.Marshal(our2)
	fmt.Println(string(convert))
	return convert
}

func decode(convert []byte) {


	res := our{}
	json.Unmarshal(convert, &res)
	fmt.Println("decode", res)

}

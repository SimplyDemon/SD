package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)


var identifier, platform, text = "aaccb", "ios", "cedax"

=======
func main() {
	var identifier, platform, text = "aaaab", "io1s", "teeec"
	encode(identifier, platform, text)
	theString := encode(identifier, platform, text)
	decode(theString)
	s := string(theString)

	var z = decode(theString)
	fmt.Println("its Z", z)

	if isJson(s) {
		fmt.Println("ITS A JSON")
	} else {
		fmt.Println(s, " NOT A JSON")
	}
}


type Request struct {
	Identifier string     `json:"identifier"`
	Platform   string       `json:"platform"`
	Text       string           `json:"text"`
}


func main() {

	encode(identifier, platform, text)
	theString := encode(identifier, platform, text)
	decode(theString)
=======
func isValidIdentifier(ID string) bool {


	matched, err := regexp.MatchString("^[A-Fa-f\\d]{5}$", ID)
	return matched || err != nil
}
func encode(identifier string, platform string, text string) []byte {

	curRequest := &Request{
		Identifier:   identifier,
		Platform:   platform,
		Text:   text}


	jsonData, _ := json.Marshal(curRequest)
	fmt.Println(string(jsonData))
	return jsonData

	convert, _ := json.Marshal(curRequest)
	fmt.Println("enc", string(convert))
	fmt.Println(len(convert))
	return convert

}
func decode(jsonData []byte) Request {


func decode(jsonData []byte) {



	res := Request{}
	json.Unmarshal(jsonData, &res)
	fmt.Println("decode", res)
	return res

}
func isJson(s string) bool {

	matched, err := regexp.MatchString("^[{\"identifier]+[\\w\\W]+[\"}]+$", s)
	return matched || err != nil
}
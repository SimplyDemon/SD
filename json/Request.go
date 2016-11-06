package main

import "encoding/json"


type Request struct {
	Identifier string     `json:"identifier"`
	Platform   string       `json:"platform"`
	Text       string           `json:"text"`
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

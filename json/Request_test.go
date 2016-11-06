package main

import "testing"

func TestDecode(t *testing.T) {
	var id [5] string
	id[0] = "aaaaa"
	id[1] = "bbbbb"
	id[2] = "ccccc"
	id[3] = "ddddd"
	id[4] = "eeeee"
	var pt [5] string
	pt[0] = "ios"
	pt[1] = "ms"
	pt[2] = "cli"
	pt[3] = "windows"
	pt[4] = "android"
	var tx [5] string
	tx[0] = "1"
	tx[1] = "2"
	tx[2] = "3"
	tx[3] = "4"
	tx[4] = "5"

	var result[5] Request
	result[0] = decode("{\"identifier\":\"" + id[0] + "\", \"platform\":\"" + pt[0] + "\", \"text\":\"" + tx[0] + "\"}")
	result[1] = decode("{\"identifier\":\"" + id[1] + "\", \"platform\":\"" + pt[1] + "\", \"text\":\"" + tx[1] + "\"}")
	result[2] = decode("{\"identifier\":\"" + id[2] + "\", \"platform\":\"" + pt[2] + "\", \"text\":\"" + tx[2] + "\"}")
	result[3] = decode("{\"identifier\":\"" + id[3] + "\", \"platform\":\"" + pt[3] + "\", \"text\":\"" + tx[3] + "\"}")
	result[4] = decode("{\"identifier\":\"" + id[4] + "\", \"platform\":\"" + pt[4] + "\", \"text\":\"" + tx[4] + "\"}")
	for i := 0; i < 5; i++ {

		if result[i].Identifier != id[i] && result[i].Platform != pt[i] && result[i].Text != tx[i] {
			t.Error("Expected == ", result)
		}
	}

	var result2[5] Request
	result2[0] = decode(id[0] + pt[0] + tx[0])
	result2[1] = decode("id[0]  + pt[0] + tx[0] + 3132132")
	result2[2] = decode("")
	result2[3] = decode(" ")
	result2[4] = decode("Identifier Platform Text")
	for i := 0; i < 5; i++ {

		if result2[i].Identifier == id[i] && result2[i].Platform == pt[i] && result2[i].Text == tx[i] {
			t.Error("Expected != ", result)
		}
	}
}


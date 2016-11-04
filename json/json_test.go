package main

import (
	"testing"

	"strings"
	"time"

	"fmt"
	"os"
	"log"
	"bufio"
	"io"
)

func TestIsValidPlatform(t *testing.T) {
	var result[5] bool
	result[0] = isValidPlatform("ios")
	result[1] = isValidPlatform("ms")
	result[2] = isValidPlatform("browser")
	result[3] = isValidPlatform("cli")
	result[4] = isValidPlatform("android")
	for i := 0; i < 5; i++ {
		if result[i] != true {
			t.Error("Expected true ", result)
		}
	}
	var result2[3] bool
	result2[0] = isValidPlatform("kenbuchi")
	result2[1] = isValidPlatform("123531")
	result2[2] = isValidPlatform("russianOS")
	for i := 0; i < 3; i++ {
		if result2[i] == true {
			t.Error("Expected false ", result)
		}
	}

}

func TestIsValidIdentifier(t *testing.T) {
	var result[5] bool
	result[0] = isValidIdentifier("aaaaa")
	result[1] = isValidIdentifier("AabcD")
	result[2] = isValidIdentifier("a1b2C")
	result[3] = isValidIdentifier("12345")
	result[4] = isValidIdentifier("ABfDE")
	for i := 0; i < 5; i++ {

		if result[i] != true {
			t.Error("Expected true ", result)
		}
	}

	var result2[5] bool
	result2[0] = isValidIdentifier("AbcDH")
	result2[1] = isValidIdentifier("1!34U")
	result2[2] = isValidIdentifier("qwertyuiop")
	result2[3] = isValidIdentifier("abcd#")
	result2[4] = isValidIdentifier("AUfjv")
	for i := 0; i < 5; i++ {

		if result2[i] == true {
			t.Error("Expected false ", result)
		}
	}

}

func TestIsAnyText(t *testing.T) {
	var result[5] bool
	result[0] = isAnyText("$#@#$%!")
	result[1] = isAnyText(" ")
	result[2] = isAnyText("QWErty")
	result[3] = isAnyText("12345dsfgsdfgsd[fio jgsdofijgosdpifgjsodpifjgpodsifjhgposdifgj 111111111111111")
	result[4] = isAnyText("/")
	for i := 0; i < 5; i++ {

		if result[i] != true {
			t.Error("Expected true ", result)
		}
	}
	var result2 bool
	result2 = isAnyText("")
	if result2 == true {
		t.Error("Expected false ", result)
	}
}

func TestIsJson(t *testing.T) {
	var result[5] bool
	result[0] = isJson("{\"$#@#$%\"}")
	result[1] = isJson("{\" \"}")
	result[2] = isJson("{\":WErty:\"}")
	result[3] = isJson("{\"12345dsfgsdfgsd[fio jgsdofijgosdpifgjsodpifjgpodsifjhgposdifgj 111111111111111\"}")
	result[4] = isJson("{\"/\"}")
	for i := 0; i < 5; i++ {

		if result[i] != true {
			t.Error("Expected true ", result)
		}
	}
	var result2[5] bool
	result2[0] = isJson("{\"AbcDH")
	result2[1] = isJson("{1!34U")
	result2[2] = isJson("\"{qwertyuiop\"{")
	result2[3] = isJson("{{abcd#!")
	result2[4] = isJson(" \"AUfjv\" }")
	for i := 0; i < 5; i++ {

		if result2[i] == true {
			t.Error("Expected false ", result)
		}
	}

}
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

func next_to_last(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	bf := bufio.NewReader(f)
	prev_line := ""

	for {
		switch line, err := bf.ReadString('\n'); err {

		case nil:
			// valid line, echo it.  note that line contains trailing \n.

			prev_line = line // Save line as previous

		case io.EOF:
			if line > "" {
				// last line of file missing \n, but still valid

			}
			return prev_line // Return previous line

		default:
			log.Fatal(err)
		}
	}
}

func TestWriteLog(t *testing.T) {

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

	for i := 0; i < 5; i++ {

		writeLog(id[i], pt[i], tx[i])
		line := next_to_last("./json/" + time.Now().Format("2006.01.02") + ".logs")

		x := strings.TrimSpace(line)
		z := strings.Split(x, ":")
		c := strings.Split(z[len(z) - 1], ".")
		c1 := strings.Split(z[len(z) - 2], ".")
		c2 := strings.Split(z[len(z) - 3], ".")
		v := strings.Split(c[0], " ")
		v1 := strings.Split(c1[0], " ")
		v2 := strings.Split(c2[0], " ")
		fmt.Println(v[1], v1[1], v2[1])

		if v[1] != tx[i]&& v1[1] != pt[i] && v2[1] != id[i] {
			t.Error("Expected == ")
		}
	}
}
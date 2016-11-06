package main

import (
	"testing"
	"time"
	"strings"
	"fmt"
	"io"
	"log"
	"os"
	"bufio"
)

func readLastLine(filePath string) string {
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
		line := readLastLine("./json/" + time.Now().Format("2006.01.02") + ".logs")

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
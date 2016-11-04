package main

import (
	"bufio"
	"os"
	"io"
	"fmt"
)

func main() {
	file, err := os.OpenFile("./json/text.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {

		fmt.Println("can't open file ")
	}
	reader := bufio.NewReader(file)

	var line string
	for {

		if err != io.EOF {
			line, err = reader.ReadString(1)

		}

		break
	}
	fmt.Println("123", line)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"os"
)

func main() {

	fmt.Println(next_to_last("json/json/2016.11.04.logs"))

}


func next_to_last(z string) string {
	f, err := os.Open(z)
	if err != nil {
		log.Fatal(err)
	}
	bf := bufio.NewReader(f)
	prev_line := ""

	for {
		switch line, err := bf.ReadString('\n'); err {

		case nil:
			// valid line, echo it.  note that line contains trailing \n.
			fmt.Println(line)
			prev_line = line // Save line as previous

		case io.EOF:
			if line > "" {
				// last line of file missing \n, but still valid
				fmt.Println(line, "in func")
			}
			return  prev_line // Return previous line

		default:
			log.Fatal(err)
		}
	}
}
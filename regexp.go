package main

import (
	"fmt"
	"os"
	"regexp"
	"flag"
	"bufio"
	"log"
	"time"
	"path/filepath"

)

func main() {

	var identifier, platform, text = getArgs()

	if isValidPlatform(platform) && len(text) > 0 && isValidIdentifier(identifier) {
		fmt.Println(identifier, platform, text)
		writeLog(identifier, platform, text)
	} else {
		printHelp()
		os.Exit(0)
	}

}

func printHelp() {
	fmt.Println("Usage logger:\n" +
		"logger <-i=your identifier> <-p=your platform> <-t=some text comments>\n" +
		"where:\n" +
		"\t-i - your identifier - unique value (latin letters and digits only, lenght = 5)\n" +
		"\t-p - your platform - chose one of: browser, ios, android, cli, ms\n" +
		"\t-t - enter some text comments.")
}

func isValidPlatform(platform string) bool {
	switch platform {
	case
		"browser", "ios", "android", "cli", "ms":
		return true
	}
	return false
}

func isValidIdentifier(identifier string) bool {

	matched, err := regexp.MatchString("^[A-Fa-f\\d]{5}$", identifier)
	return matched || err != nil
}

func getArgs() (identifier string, platform string, text string) {
	if len(os.Args) != 4 {
		identifier = ""
		platform = ""
		text = ""

	} else {
		flag.StringVar(&identifier, "i", "", "Enter your identifier correctly(-i), must contain latin letters and digits only, lenght = 5")
		flag.StringVar(&platform, "p", "", "Enter your platform correctly(-p) - one of: browser, ios, android, cli, ms")
		flag.StringVar(&text, "t", "", "Enter some text")

		flag.Parse()
	}
	return identifier, platform, text
}

func writeLog(identifier string, platform string, text string) {
	folderPath := filepath.Join(identifier, platform)
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)

	}
	filePath := filepath.Join(identifier, platform, time.Now().Format("2006.01.02"))
	file, err := os.OpenFile(filePath + ".log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)


	fmt.Fprintln(writer,"\n" + "Identifier is:" + identifier + ". Platform  is:" + platform + ". Text is: " + text + ". Time and date is:" + time.Now().Format("2006.01.02.15:04:05"))
	writer.Flush()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	defer file.Close()


}
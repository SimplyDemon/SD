package main

import (
	"os"
	"log"
	"time"
)



func LogText(funcName string) {
	f, err := os.OpenFile("./debugFile.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(funcName)
}

func writeLog(identifier string, platform string, text string) {

	f, err := os.OpenFile("./json/" + time.Now().Format("2006.01.02") + ".logs", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Print("\n" + "Identifier is: " + identifier + ". Platform  is: " + platform + ". Text is: " + text + "." + "\n")
}
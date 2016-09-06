package main

import (

	"testing"
	"path/filepath"
	"time"
	"os"
	"bufio"
	"io"
	"fmt"
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

func TestWriteLog(t *testing.T) {

	identifier := "Abbbb"
	platform := "ios"
	text := "some text"

	if isValidIdentifier(identifier) != true {
		t.Error("Invalid Identifier!")
	}
	if isValidPlatform(platform) != true {
		t.Error("Invalid Platform!")
	}
	if len(text) == 0 {
		t.Error("Enter some text")
	}

	writeLog(identifier, platform, text)       // вызывает функцию writeLog с указаными аргументами
	filePath := filepath.Join("./logs", identifier, platform, time.Now().Format("2006.01.02"))
	file, err := os.OpenFile(filePath + ".log", os.O_APPEND, 0644)

	if err != nil {
		// если не удаётся открыть созданный файл - вывести сообщение об ошибке
		t.Error("can't open ")
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString(1)
		if err == io.EOF {
			fmt.Println(line)
			break

		} else if err != nil {

			t.Error("error! ")             // if you return error

		}

	}

	defer file.Close()

}

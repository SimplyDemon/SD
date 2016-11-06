package main

import (
	"strconv"
	"fmt"
	"os"
	"regexp"
	"strings"
)


func isAnyText(text string) bool {
	LogText("DEBUG: func isAnyText with arguments " + text + " started!")
	LogText("DEBUG: func isAnyText is : " + strconv.FormatBool(len(text) > 0))
	return len(text) > 0
}
func isValidPlatform(platform string) bool {
	LogText("DEBUG: func isValidPlatform with arguments " + platform + " started!")
	switch platform {
	case
		"browser", "ios", "android", "cli", "ms":
		LogText("DEBUG: func isValidPlatform is: true")
		return true
	}

	LogText("DEBUG: func isValidPlatform is : false")
	return false
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isValidIdentifier(ID string) bool {
	LogText("DEBUG: func isValidIdentifier with arguments " + ID + " started!")
	matched, err := regexp.MatchString("^[A-Fa-f\\d]{5}$", ID)
	LogText("DEBUG: func isValidIdentifier is : " + strconv.FormatBool(matched))
	return matched || err != nil
}


func isJson(s string) bool {
	LogText("DEBUG: func isJson with arguments " + s + " started!")
	trim := strings.TrimSpace(s)
	check := string(trim[0]) == "{" && string(trim[len(trim) - 1]) == "}"

	LogText("DEBUG: func isJson is : " + strconv.FormatBool(check))
	return check
}
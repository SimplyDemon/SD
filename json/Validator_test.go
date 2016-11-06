package main

import "testing"

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

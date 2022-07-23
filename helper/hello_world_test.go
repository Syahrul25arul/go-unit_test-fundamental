package helper

import "testing"

func TestHelloWorldSucces(t *testing.T) {
	result := HelloWorld("Hendrik")
	if result != "Hello Hendrik" {
		panic("Result is not Hendrik")
	}
}

func TestHelloWorldFailed(t *testing.T) {
	result := HelloWorld("Test")
	if result != "Hello Hendrik" {
		panic("Result is not Hendrik")
	}
}

package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldSucces(t *testing.T) {
	result := HelloWorld("Hendrik")
	if result != "Hello Hendrik" {
		// error
		// t.Fail() // tidak menghentikan function unit test
		t.Error("Result must be 'Hello Hendrik'") // akan memunculkan pesan dan memanggil function t.Fail()
	}

	fmt.Println("Test Hello World Success")
}

func TestHelloWorldFailed(t *testing.T) {
	result := HelloWorld("Hendrik")
	if result != "Hello Hendrik" {
		// t.FailNow() // menghentikan function unit test
		t.Fatal("result must be 'Hello Hendrik'") // Akan memunculkan pesan dan memanggil func t.FailNow()
	}

	fmt.Println("Test Hello World Failed")
}

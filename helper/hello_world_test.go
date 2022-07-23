package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldSucces(t *testing.T) {
	assert := assert.New(t)
	result := HelloWorld("Hendrik")

	assert.Equal(result, "Hello Hendrik", "Message must be Hello Hendrik") // ini sama seperti t.Error() yang memanggil t.Fail()
	fmt.Println("TestHelloWorldSucces with assert equal")
}

func TestHelloWorldFailed(t *testing.T) {
	result := HelloWorld("Jamal")
	require := require.New(t)

	require.Equal(result, "Hello Hendrik", "Result must be Hello Hendrik") // ini sama seperti t.Fatal() yang memanggil t.FailNow()

	fmt.Println("Test Hello World Failed")
}

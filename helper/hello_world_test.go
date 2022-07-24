package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After unit Test")
	// after
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run Windows")
	}
	result := HelloWorld("Hendrik")
	require.Equal(t, "Hello Hendrik", result, "Result must be 'Hello Hendrik'")
}

func TestHelloWorldSucces(t *testing.T) {
	assert := assert.New(t)
	result := HelloWorld("Hendrik")

	assert.Equal(result, "Hello Hendrik", "Message must be Hello Hendrik") // ini sama seperti t.Error() yang memanggil t.Fail()
	fmt.Println("TestHelloWorldSucces with assert equal")
}

func TestHelloWorldFailed(t *testing.T) {
	result := HelloWorld("Jamal")
	require := require.New(t)

	require.NotEqual(result, "Hello Hendrik", "Result must be not Hello Hendrik") // ini sama seperti t.Fatal() yang memanggil t.FailNow()

	fmt.Println("Test Hello World Failed")
}

func TestSubTest(t *testing.T) {
	t.Run("Hendrik", func(t *testing.T) {
		result := HelloWorld("Hendrik")

		assert.Equal(t, result, "Hai Hendrik", "Message must be Hello Hendrik")
	})
	t.Run("Array", func(t *testing.T) {
		result := HelloWorld("Jamal")

		require.NotEqual(t, result, "Hello Hendrik", "Result must be not Hello Hendrik")
	})
}

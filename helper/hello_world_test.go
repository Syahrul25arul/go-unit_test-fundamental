package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var Assert *assert.Assertions
var Require *require.Assertions

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Hendrik",
			request: "Hendrik Rizal Array",
		},
		{
			name:    "Jamal",
			request: "Jamal Adrianto",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("Hendrik", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hendrik")
		}
	})
	b.Run("Jamal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jamal")
		}
	})
}

func BenchmarkHelloWorldHendrik(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hendrik")
	}
}

func BenchmarkHelloWorldJamal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Jamal")
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")
	Assert = assert.New(&testing.T{})
	Require = require.New(&testing.T{})
	m.Run()

	fmt.Println("After unit Test")
	// after
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run Windows")
	}
	result := HelloWorld("Hendrik")
	Require.Equal("Hello Hendrik", result, "Result must be 'Hello Hendrik'")
}

func TestHelloWorldSucces(t *testing.T) {
	result := HelloWorld("Hendrik")

	Assert.Equal(result, "Hello Hendrik", "Message must be Hello Hendrik") // ini sama seperti t.Error() yang memanggil t.Fail()
	fmt.Println("TestHelloWorldSucces with assert equal")
}

func TestHelloWorldFailed(t *testing.T) {
	result := HelloWorld("Jamal")

	Require.NotEqual(result, "Hello Hendrik", "Result must be not Hello Hendrik") // ini sama seperti t.Fatal() yang memanggil t.FailNow()

	fmt.Println("Test Hello World Failed")
}

func TestSubTest(t *testing.T) {
	t.Run("Hendrik", func(t *testing.T) {
		result := HelloWorld("Hendrik")

		Assert.Equal(result, "Hai Hendrik", "Message must be Hello Hendrik")
	})
	t.Run("Array", func(t *testing.T) {
		result := HelloWorld("Jamal")

		Require.NotEqual(result, "Hello Hendrik", "Result must be not Hello Hendrik")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hendrik",
			request:  "Hendrik",
			expected: "Hello Hendrik",
		},
		{
			name:     "Rizal",
			request:  "Rizal",
			expected: "Hello Rizal",
		},
		{
			name:     "Array",
			request:  "Array",
			expected: "Hello Array",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			Assert.Equal(test.expected, result, "Result must be "+test.expected)
		})
	}
}

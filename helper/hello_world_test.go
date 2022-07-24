package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Test")
	m.Run()
	fmt.Println("After Test")
}

//Sub test
func TestSubTest(t *testing.T) {
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		assert.Equal(t, "Hello Budi", result, "Result must be Hello Budi")
	})
	t.Run("Setiawan", func(t *testing.T) {
		result := HelloWorld("Setiawan")
		assert.Equal(t, "Hello Setiawan", result, "Result must be Hello Setiawan")
	})
}

// Table Test
func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Setiawan",
			request:  "Setiawan",
			expected: "Hello Setiawan",
		},
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Result must be "+test.expected)
		})

	}

}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Budi")

	if result != "Hello Budi" {
		// panic("Testing Gagal")
		// t.Fail()
		t.Error("Result must be Hello Budi")
	}
	// fmt.Println("Masih di Eksekusi")

}
func TestHelloWorldSetiawan(t *testing.T) {
	result := HelloWorld("Setiawan")

	if result != "Hello Setiawan" {
		// panic("Testing Gagal")
		// t.FailNow()
		t.Fatal("Result must be Hello Setiawan")
	}

	fmt.Println("ini tidak dieksekusi")

}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Budi")
	// Pakai t.Fail()
	assert.Equal(t, "Hello Budi", result)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Budi")
	// Pakai t.FailNow()
	require.Equal(t, "Hello Budi", result)
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skip Windows")
	}
	result := HelloWorld("Budi")
	// Pakai t.FailNow()
	require.Equal(t, "Hello Budi", result, "result must be Hello Budi")
}

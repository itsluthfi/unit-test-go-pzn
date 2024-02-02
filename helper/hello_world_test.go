package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -v -run=NoTest -bench=(nama benchmark)
// go test -v -run=NoTest -bench=(nama benchmark)/(nama subtest)
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Luthfi",
			request: "Luthfi",
		},
		{
			name:    "Izzuddin",
			request: "Izzuddin",
		},
		{
			name:    "LuthfiIzzuddin",
			request: "Luthfi Izzuddin",
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

func BenchmarkSub(b *testing.B) {
	b.Run("Luthfi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Luthfi")
		}
	})
	b.Run("Izzuddin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Izzuddin")
		}
	})
}

func BenchmarkHelloWorldLuthfi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Luthfi")
	}
}

func BenchmarkHelloWorldIzzuddin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Izzuddin")
	}
}

// go test -v (run semua test yang ada di package)
// go test -v -run=(nama test)
// go test -v -run=(nama test)/(nama subtest)
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE TEST")

	m.Run()

	// after
	fmt.Println("AFTER TEST")
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Luthfi",
			request:  "Luthfi",
			expected: "Hello Luthfi",
		},
		{
			name:     "Izzuddin",
			request:  "Izzuddin",
			expected: "Hello Izzuddin",
		},
		{
			name:     "Hanif",
			request:  "Hanif",
			expected: "Hi Hanif",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Luthfi", func(t *testing.T) {
		result := HelloWorld("Luthfi")
		require.Equal(t, "Hello Luthfi", result, "Result must be 'Hello Luthfi'")
	})

	t.Run("Izzuddin", func(t *testing.T) {
		result := HelloWorld("Izzuddin")
		require.Equal(t, "Hi Izzuddin", result, "Result must be 'Hello Izzuddin'")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run test on Windows")
	}

	// test gaakan dijalankan, bakal ada log SKIP dan pesan kayak di atas
	result := HelloWorld("Luthfi")
	require.Equal(t, "Hello Luthfi", result, "Result must be 'Hello Luthfi'")
}

// direkomendasikan kalo unit testing pakenya assert/require daripada if-else
func TestHelloWOrldRequire(t *testing.T) {
	result := HelloWorld("Luthfi")
	require.Equal(t, "Hello Luthfi", result, "Result must be 'Hello Luthfi'")
	// require, dia bakal manggil t.FailNow()
	fmt.Println("Test done") // ga keprint
}

func TestHelloWOrldAssertion(t *testing.T) {
	result := HelloWorld("Luthfi")
	assert.Equal(t, "Hello Luthfi", result, "Result must be 'Hello Luthfi'")
	// kalo gagal bakal manggil t.Fail()
	// ada juga require, tapi kalo gagal dia bakal manggil t.FailNow()
	fmt.Println("Test done") // bakal keprint
}

func TestHelloWorldLuthfi(t *testing.T) {
	result := HelloWorld("Luthfi")
	if result != "Hello Luthfi" {
		// t.Fail()
		// dia bakal ngasih fail tapi kode program di bawahnya tetep dijalanin
		t.Error("Result must be 'Hello Luthfi'") // ngasih log stringnya dulu terus baru eksekusi t.Fail()
	}

	fmt.Println("Test done") // bakal keprint
}

func TestHelloWorldIzzuddin(t *testing.T) {
	result := HelloWorld("Izzuddin")
	if result != "Hello Izzuddin" {
		// t.FailNow()
		// dia bakal ngasih fail tapi kode program di bawahnya ga dijalanin/langsung berhenti pas fail
		t.Fatal("Result must be 'Hello Luthfi'") // ngasih log stringnya dulu terus baru eksekusi t.FailNow()
	}

	fmt.Println("Test done") // ga keprint
}

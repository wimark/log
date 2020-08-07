package log

import (
	"log"
	"os"
	"testing"
)

func BenchmarkInfo(b *testing.B) {
	defer quiet()()
	InitSingleStr("1")

	for i := 0; i < b.N; i++ {
		Info("i: %d", i)
	}
}

func BenchmarkDebug(b *testing.B) {
	defer quiet()()
	os.Setenv("LOGLEVEL", "DEBUG")
	InitSingleStr("2")

	for i := 0; i < b.N; i++ {
		Debug("i: %d", i)
	}
}

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

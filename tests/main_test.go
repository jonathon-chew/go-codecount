package main_test

import (
	"os"
	"testing"
)

func Benchmark_Speed(b *testing.B) {

	os.Chdir("/Users/hunteradder626/Documents/Scripts")

	os.Args = append(os.Args, "-i", ".git", "venv")

	for b.Loop() {
		main()
	}
}

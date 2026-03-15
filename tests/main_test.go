package main_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func Benchmark_Speed(b *testing.B) {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		b.Fatal("unable to determine test file location")
	}

	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(filename), "..", "..", "..", ".."))
	if err := os.Chdir(repoRoot); err != nil {
		b.Fatalf("unable to change directory to repo root: %v", err)
	}

	os.Args = append(os.Args, "-i", ".git", "venv")

	for b.Loop() {
		main()
	}
}

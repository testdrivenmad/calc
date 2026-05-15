package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// binPath holds the path to the compiled calc binary, shared across tests.
var binPath string

// TestMain builds the calc binary once before running the test cases, then
// removes it afterward. Building once avoids paying `go build` cost per test.
func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "calc-test-")
	if err != nil {
		panic(err)
	}
	binPath = filepath.Join(dir, "calc")
	build := exec.Command("go", "build", "-o", binPath, ".")
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		os.RemoveAll(dir)
		panic(err)
	}

	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}

// run invokes the compiled binary with args and returns stdout, stderr, and
// the process exit code.
func run(t *testing.T, args ...string) (string, string, int) {
	t.Helper()
	cmd := exec.Command(binPath, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	code := 0
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			code = ee.ExitCode()
		} else {
			t.Fatalf("unexpected exec error: %v", err)
		}
	}
	return stdout.String(), stderr.String(), code
}

func TestCalc(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantOut  string // expected stdout (trimmed of trailing newline)
		wantCode int
		wantErr  string // substring expected in stderr (empty = no check)
	}{
		{
			name:     "add two ints",
			args:     []string{"add", "1", "2"},
			wantOut:  "3",
			wantCode: 0,
		},
		{
			name:     "add symbol alias",
			args:     []string{"+", "1", "2", "3", "4"},
			wantOut:  "10",
			wantCode: 0,
		},
		{
			name:     "sub multiple",
			args:     []string{"sub", "10", "1", "2", "3"},
			wantOut:  "4",
			wantCode: 0,
		},
		{
			name:     "sub symbol alias",
			args:     []string{"-", "5", "2"},
			wantOut:  "3",
			wantCode: 0,
		},
		{
			name:     "floats add",
			args:     []string{"add", "0.1", "0.2", "0.3"},
			wantOut:  "0.6000000000000001", // floating-point reality from float64 addition
			wantCode: 0,
		},
		{
			name:     "negative result",
			args:     []string{"sub", "1", "5"},
			wantOut:  "-4",
			wantCode: 0,
		},
		{
			name:     "too few args prints usage",
			args:     []string{"add", "1"},
			wantCode: 2,
			wantErr:  "usage:",
		},
		{
			name:     "no args prints usage",
			args:     []string{},
			wantCode: 2,
			wantErr:  "usage:",
		},
		{
			name:     "unknown operator prints usage",
			args:     []string{"mul", "2", "3"},
			wantCode: 2,
			wantErr:  "usage:",
		},
		{
			name:     "invalid number",
			args:     []string{"add", "1", "abc"},
			wantCode: 2,
			wantErr:  `invalid number "abc"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, code := run(t, tc.args...)
			if code != tc.wantCode {
				t.Errorf("exit code = %d, want %d\nstderr: %s", code, tc.wantCode, stderr)
			}
			gotOut := strings.TrimRight(stdout, "\n")
			if tc.wantOut != "" && gotOut != tc.wantOut {
				t.Errorf("stdout = %q, want %q", gotOut, tc.wantOut)
			}
			if tc.wantErr != "" && !strings.Contains(stderr, tc.wantErr) {
				t.Errorf("stderr = %q, want substring %q", stderr, tc.wantErr)
			}
		})
	}
}

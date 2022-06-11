package aws

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

var testStdoutValue = "testing"

func TestVersionSuccess(t *testing.T) {
	var bytes []byte
	var err error

	bytes, err = Version(newFakeExecCommand(false))

	assert.Nil(t, err, "Version should not error when aws exists")

	stringResponse := string(bytes)
	assert.EqualValues(t, testStdoutValue, stringResponse, "Expected executable to return my output")
}

func TestRunningCommandFailure(t *testing.T) {
	var bytes []byte
	var err error

	bytes, err = Version(newFakeExecCommand(true))
	assert.NotNil(t, err, "Expecting command to fail")
	assert.NotNilf(t, bytes, "Reader should still be set in error condition")

	output := string(bytes)
	assert.EqualValues(t, testStdoutValue, output)
}

// TestShellProcessSuccess is a method that is called as a substitute for a shell command,
// the GO_TEST_PROCESS flag ensures that if it is called as part of the test suite, it is
// skipped.
func TestShellProcess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}

	_, err := fmt.Fprintf(os.Stdout, testStdoutValue)
	if err != nil {
		os.Exit(1)
	}

	if os.Getenv("GO_TEST_PROCESS") == "2" {
		os.Exit(1)
	}

	exitIntCode, err := strconv.Atoi(os.Getenv("GO_TEST_EXIT"))
	if err != nil {
		log.Fatalf("TEST ENV NOT SETUP PROPERLY")
		return
	}

	os.Exit(exitIntCode)
}

func newFakeExecCommand(shouldFail bool) func(string, ...string) *exec.Cmd {
	var exitCode = 0
	if shouldFail {
		exitCode = 1
	}
	return func(command string, args ...string) *exec.Cmd {
		cs := []string{"-test.run=TestShellProcess", "--", command}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_TEST_PROCESS=1", fmt.Sprintf("GO_TEST_EXIT=%d", exitCode)}
		return cmd
	}
}

package bugapp

import (
	"fmt"
	"github.com/driusan/bug/bugs"
	//	"io"
	"io/ioutil"
	"os"
	"testing"
)

func captureOutput(f func(), t *testing.T) (string, string) {
	// Capture STDOUT with a pipe
	stdout := os.Stdout
	stderr := os.Stderr
	so, op, _ := os.Pipe() //outpipe
	oe, ep, _ := os.Pipe() //errpipe
	defer func(stdout, stderr *os.File) {
		os.Stdout = stdout
		os.Stderr = stderr
	}(stdout, stderr)

	os.Stdout = op
	os.Stderr = ep

	f()

	os.Stdout = stdout
	os.Stderr = stderr

	op.Close()
	ep.Close()

	errOutput, err := ioutil.ReadAll(oe)
	if err != nil {
		t.Error("Could not get output from stderr")
	}
	stdOutput, err := ioutil.ReadAll(so)
	if err != nil {
		t.Error("Could not get output from stdout")
	}
	return string(stdOutput), string(errOutput)
}

// Captures stdout and stderr to ensure that
// a usage line gets printed to Stderr when
// no parameters are specified
func TestCreateHelpOutput(t *testing.T) {
	config := bugs.Config{}
	stdout, stderr := captureOutput(func() {
		Create(ArgumentList{}, config)
	}, t)

	if stdout != "" {
		t.Error("Unexpected output on stdout.")
	}
	if stderr[:7] != "Usage: " {
		t.Error("Expected usage information with no arguments")
	}

}

// Test a very basic invocation of "Create" with the -n
// argument. We can't try it without -n, since it means
// an editor will be spawned..
func TestCreateNoEditor(t *testing.T) {
	config := bugs.Config{}
	dir, err := ioutil.TempDir("", "createtest")
	if err != nil {
		t.Error("Could not create temporary dir for test")
		return
	}
	os.Chdir(dir)
	os.MkdirAll("issues", 0700)
	defer os.RemoveAll(dir)
	// On MacOS, /tmp is a symlink, which causes GetDirectory() to return
	// a different path than expected in these tests, so make the issues
	// directory explicit with an environment variable
	err = os.Setenv("PMIT", dir)
	if err != nil {
		t.Error("Could not set environment variable: " + err.Error())
		return
	}

	stdout, stderr := captureOutput(func() {
		Create(ArgumentList{"-n", "Test", "bug"}, config)
	}, t)
	if stderr != "" {
		t.Error("Unexpected output on STDERR for Test-bug")
	}
	if stdout != "Created issue: Test bug\n" {
		t.Error("Unexpected output on STDOUT for Test-bug")
	}
	issuesDir, err := ioutil.ReadDir(fmt.Sprintf("%s/issues/", dir))
	if err != nil {
		t.Error("Could not read issues directory")
		return
	}
	if len(issuesDir) != 1 {
		t.Error("Unexpected number of issues in issues dir\n")
	}

	bugDir, err := ioutil.ReadDir(fmt.Sprintf("%s/issues/Test-bug", dir))
	if len(bugDir) != 1 {
		t.Error("Unexpected number of files found in Test-bug dir\n")
	}
	if err != nil {
		t.Error("Could not read Test-bug directory")
		return
	}

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/issues/Test-bug/Description", dir))
	if err != nil {
		t.Error("Could not load description file for Test bug" + err.Error())
	}
	if len(file) != 0 {
		t.Error("Expected empty file for Test bug")
	}
}

/* currently hangs spawning editor

// this test will not spawn editor
func TestCreateNoIssuesDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "createtest")
	if err != nil {
		t.Error("Could not create temporary dir for test")
		return
	}
	os.Chdir(dir)
	// This is what we are testing for
	//os.MkdirAll("issues", 0700)
	//defer os.RemoveAll(dir)
	err = os.Setenv("PMIT", string(dir))
	if err != nil {
		t.Error("Could not set environment variable: " + err.Error())
		return
	}

	stdout, stderr := captureOutput(func() {
		Create(ArgumentList{"Test", "bug"}) // fire editor without -n
	}, t)
	_ = stderr
	_ = stdout
	// stderr is expected from log.Fatal(err)
	//if stdout != "" {
	//	t.Error("Unexpected output on STDOUT for Test-bug")
	//}
}
*/

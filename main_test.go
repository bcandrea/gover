package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGetVersion(t *testing.T) {
	v, err := GetVersion("github.com/bcandrea/gover/test")
	if err != nil {
		t.Fatal(err)
	}
	if expected, got := "0.4.2", v; expected != got {
		t.Errorf("Expected version %s, got %s", expected, got)
	}
}

func TestGetError(t *testing.T) {
	_, err := GetVersion("github.com/bcandrea/gover/noversion")
	if err == nil {
		t.Error("Trying to get the version from a package that does not have one should return an error")
	}
	_, err = GetVersion("github.com/bcandrea/thispackagedoesnotexist")
	if err == nil {
		t.Error("Trying to get the version from a package that does not exist should return an error")
	}
}

func TestGetVersionVar(t *testing.T) {
	v, err := GetVersion("github.com/bcandrea/gover/testvar")
	if err != nil {
		t.Fatal(err)
	}
	if expected, got := "0.9.9", v; expected != got {
		t.Errorf("Expected version %s, got %s", expected, got)
	}
}

func TestGetVersionInFunc(t *testing.T) {
	_, err := GetVersion("github.com/bcandrea/gover/testfunc")
	if err == nil {
		t.Error("Version variables defined in functions should not be considered")
	}
}

func TestGetVersionNotLiteral(t *testing.T) {
	_, err := GetVersion("github.com/bcandrea/gover/nostring")
	if err == nil {
		t.Error("Versions should be basic literals")
	}
}

func TestGetVersionInt(t *testing.T) {
	v, err := GetVersion("github.com/bcandrea/gover/testint")
	if err != nil {
		t.Fatal(err)
	}
	if expected, got := "12", v; expected != got {
		t.Errorf("Expected version %s, got %s", expected, got)
	}
}

func TestVersionInMain(t *testing.T) {
	v, err := GetVersion("github.com/bcandrea/gover")
	if err != nil {
		t.Fatal(err)
	}
	if expected, got := Version, v; expected != got {
		t.Errorf("Expected version %s, got %s", expected, got)
	}
}

func TestRelativePath(t *testing.T) {
	v, err := GetVersion("./test")
	if err != nil {
		t.Fatal(err)
	}
	if expected, got := "0.4.2", v; expected != got {
		t.Errorf("Expected version %s, got %s", expected, got)
	}

	v, err = GetVersion(".")
	if err != nil {
		t.Fatal(err)
	}
	if expected, got := Version, v; expected != got {
		t.Errorf("Expected version %s, got %s", expected, got)
	}
}

func TestGetVersionWithNoPackage(t *testing.T) {
	_, err := GetVersion("github.com/bcandrea/gover/data")
	if err == nil {
		t.Error("Trying to get the version from a package that does not contain go files should return an error")
	}
}

func TestGetVersionOfFile(t *testing.T) {
	_, err := GetVersion("./test/version.go")
	if err == nil {
		t.Error("Trying to get the version from a .go file should return an error")
	}
}

func TestRun(t *testing.T) {
	var outW, errW bytes.Buffer
	args := []string{"gover", "-v"}

	if expected, got := 0, run(args, &outW, &errW); expected != got {
		t.Errorf("`gover -v' should exit with %d, got %d", expected, got)
	}

	if expected, got := fmt.Sprintf("gover version %s\n", Version), outW.String(); expected != got {
		t.Errorf("`gover -v' should return [%s], got [%s]", expected, got)
	}
}

func TestRunOnPackage(t *testing.T) {
	var outW, errW bytes.Buffer
	args := []string{"gover", "github.com/bcandrea/gover/test"}

	if expected, got := 0, run(args, &outW, &errW); expected != got {
		t.Errorf("expected exit code %d, got %d", expected, got)
	}

	if expected, got := "0.4.2\n", outW.String(); expected != got {
		t.Errorf("version should be [%s], got [%s]", expected, got)
	}
}

func TestRunWithNoArgs(t *testing.T) {
	var outW, errW bytes.Buffer
	args := []string{"gover"}

	if expected, got := 1, run(args, &outW, &errW); expected != got {
		t.Errorf("`gover' should exit with %d, got %d", expected, got)
	}
}

func TestRunWithError(t *testing.T) {
	var outW, errW bytes.Buffer
	args := []string{"gover", "github.com/bcandrea/thispackagedoesnotexist"}

	if expected, got := 2, run(args, &outW, &errW); expected != got {
		t.Errorf("expected exit code %d, got %d", expected, got)
	}
}

package main

import "testing"

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

package util

import (
	"os"
	"testing"
)

var SOME_DEFAULT = "kdfnsdlkfj23523sdgdfh"
var dummyEnv = "SOME_DUMMY"
var SOME_VAL = "aaasas"

func TestGetenvExist(t *testing.T) {
	os.Setenv(dummyEnv, SOME_VAL)
	res := Getenv(dummyEnv, SOME_DEFAULT)

	if res == "" {
		t.Error("Existing env var exist, should not return empty string")
	}
	if res == SOME_DEFAULT {
		t.Error("Existing env var exist, should not return the default")
	}
}

func TestGetenvNotExistExpectDefault(t *testing.T) {
	res := Getenv(SOME_DEFAULT, SOME_DEFAULT)

	if res == "" {
		t.Error("var does not exist, should return default")
	}
	if res != SOME_DEFAULT {
		t.Error("var does not exist, should return default")
	}
}

func TestGetenvNoDefault(t *testing.T) {
	os.Setenv(dummyEnv, SOME_VAL)
	res := GetenvNoDefault(dummyEnv)

	if res == "" {
		t.Error("Existing env var exist, should not return empty string")
	}
}

func TestGetenvNoDefaultWithNotExistingEnv(t *testing.T) {
	res := GetenvNoDefault(SOME_DEFAULT)

	if res != "" {
		t.Error("Env var does not exist, should return empty string")
	}
}

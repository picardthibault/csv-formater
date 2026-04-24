package main

import (
	"testing"
)

func TestIsFileFlagValidWithEmptyString(t *testing.T) {
	want := "--file option is mandatory !"
	err := isFileFlagValid("")
	if err == nil {
		t.Errorf(`isInputParamValid("") should return an error.`)
		return
	}
	if err.Error() != want {
		t.Errorf(`isInputParamValid("") = "%v", want match for %q`, err, want)
	}
}

func TestIsFileFlagValidWithDefinedString(t *testing.T) {
	err := isFileFlagValid("defined string")
	if err != nil {
		t.Errorf(`isInputParamValid("defined string") = "%v", want match for nil`, err)
	}
}

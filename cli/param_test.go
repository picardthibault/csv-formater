package main

import (
	"testing"
)

func TestIsInputParamValidWithEmptyString(t *testing.T) {
	want := "--input option is mandatory"
	err := isInputParamValid("")
	if err == nil {
		t.Errorf(`isInputParamValid("") should return an error.`)
		return
	}
	if err.Error() != want {
		t.Errorf(`isInputParamValid("") = "%v", want match for %q`, err, want)
	}
}

func TestIsInputParamValidWithDefinedString(t *testing.T) {
	err := isInputParamValid("defined string")
	if err != nil {
		t.Errorf(`isInputParamValid("defined string") = "%v", want match for nil`, err)
	}
}

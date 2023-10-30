package tests

import (
	"testing"

	"jguths.com/govel-panel/internal/validator"
)

func TestAddError(t *testing.T) {
	v := validator.New()

	v.AddError("test error", "test error")

	if len(v.Errors) == 0 {
		t.Errorf("add error must increment validator errors, got 0 errors added")
	}
}

func TestCheckError(t *testing.T) {
	v := validator.New()

	v.Check(true, "test", "test")
	v.Check(false, "test2", "test2")

	if v.Valid() {
		t.Errorf("check error must add at least 1 error to validator")
	}
}

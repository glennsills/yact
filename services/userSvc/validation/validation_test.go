package validation

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	const (
		arg  = "beaver@go.com" // (9 characters)
		want = true
	)

	got := IsValidEmail(arg)

	if got != want {
		t.Error("Unexpected Result")
	}

	got = IsValidEmail("")
	if got == want {
		t.Error("Unexpected Result")
	}

}

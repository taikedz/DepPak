package cli

import (
	"testing"
)

func Test_checkTrailingFlags(t *testing.T) {
	if token, ok := checkTrailingFlags([]string{"a", "b", "c"}) ; !ok {
		t.Errorf("Should have passed, got '%s'\n", token)
	}

	if _, ok := checkTrailingFlags([]string{"a", "-b", "c"}) ; ok {
		t.Errorf("Should have found token, found none.\n")
	}
}
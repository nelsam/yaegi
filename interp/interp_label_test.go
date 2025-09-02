package interp_test

import (
	"testing"

	"github.com/traefik/yaegi/interp"
)

func TestShadowedLabel(t *testing.T) {
	i := interp.New(interp.Options{})
	_, err := i.Compile(`
		package foo

		func foo() {
		someLabel:
			for {
				someLabel := "foo"
				break someLabel
			}
		}
	`)
	if err != nil {
		t.Fatalf("labels should be accessible even if shadowed by variables, but got error: %v", err)
	}
}

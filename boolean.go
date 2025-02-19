package jsonassert

import (
	"errors"
	"fmt"
)

func extractBoolean(b string) (bool, error) {
	if b == "true" {
		return true, nil
	}
	if b == "false" {
		return false, nil
	}
	return false, fmt.Errorf("could not parse '%s' as a boolean", b)
}

func (a *Asserter) checkBoolean(path string, act, exp bool) {
	a.tt.Helper()
	if act != exp {
		a.tt.Errorf("expected boolean at '%s' to be %v but was %v", path, exp, act)
	}
}

func checkBoolean(path string, act, exp bool) error {
	if act != exp {
		return errors.New(fmt.Sprintf("expected boolean at '%s' to be %v but was %v", path, exp, act))
	}
	return nil
}

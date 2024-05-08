package test

import (
	"github.com/bfun/nesbconvertpin"
	"strings"
	"testing"
)

func TestCSMP_PIN_ELEM(t *testing.T) {
	expects := []string{"_SVR", "_CLT"}
	dtas := nesbconvertpin.CSMP_PIN_ELEM()
	for _, expect := range expects {
		ok := false
		for dta, _ := range dtas {
			if strings.Contains(dta, expect) {
				ok = true
			}
		}
		if !ok {
			t.Errorf("not contains %s\n", expect)
		}
	}
}

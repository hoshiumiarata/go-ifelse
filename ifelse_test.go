package ifelse_test

import (
	"testing"

	"github.com/hoshiumiarata/go-ifelse"
)

func TestIf(t *testing.T) {
	if ifelse.If(true, 1, 2) != 1 {
		t.Error("If(true, 1, 2) should be 1")
	}
	if ifelse.If(false, 1, 2) != 2 {
		t.Error("If(false, 1, 2) should be 2")
	}
}

func TestLazyIf(t *testing.T) {
	result := ifelse.LazyIf(true, func() int {
		return 1
	}, func() int {
		t.Error("LazyIf should not call else function when condition is true")
		return 2
	})
	if result != 1 {
		t.Error("Expected LazyIf to be 1")
	}

	result = ifelse.LazyIf(false, func() int {
		t.Error("LazyIf should not call then function when condition is false")
		return 1
	}, func() int {
		return 2
	})
	if result != 2 {
		t.Error("Expected LazyIf to be 2")
	}
}

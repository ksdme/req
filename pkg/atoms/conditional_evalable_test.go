package atoms_test

import (
	"testing"

	"github.com/ksdme/req/pkg/atoms"
)

func TestTruthyConditionalValue(t *testing.T) {
	evalable := atoms.NewConditionalValue(
		atoms.NewLeafValue(1),
		atoms.NewLeafValue("a"),
		atoms.NewLeafValue("b"),
	)
	value := evalable.Evaluate(atoms.EmptyContext())

	if value != "a" {
		t.Errorf("Evaluated value and expected value did not match")
	}
}

func TestFalsyConditionalValue(t *testing.T) {
	evalable := atoms.NewConditionalValue(
		atoms.NewLeafValue(false),
		atoms.NewLeafValue("a"),
		atoms.NewLeafValue("b"),
	)
	value := evalable.Evaluate(atoms.EmptyContext())

	if value != "b" {
		t.Error("Evaluated value and expected value did not match")
	}

	// Test with the else branch being nil.
	evalable = atoms.NewConditionalValue(
		atoms.NewLeafValue(false),
		atoms.NewLeafValue("a"),
		nil,
	)
	value = evalable.Evaluate(atoms.EmptyContext())

	if value != nil {
		t.Error("Evaluated value and expected value did not match")
	}
}

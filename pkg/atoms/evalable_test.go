package atoms_test

import (
	"testing"

	"github.com/ksdme/req/pkg/atoms"
)

func context() *atoms.Context {
	return &atoms.Context{
		Variables: make(map[string]interface{}),
	}
}

func TestStringLeafValue(t *testing.T) {
	var value interface{} = "string"

	evalable := atoms.NewLeafValue(value)
	evaluated := evalable.Evaluate(context())

	if evaluated != value {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestIntLeafValue(t *testing.T) {
	var value interface{} = 1.23

	evalable := atoms.NewLeafValue(value)
	evaluated := evalable.Evaluate(context())

	if evaluated != value {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestBoolLeafValue(t *testing.T) {
	var value interface{} = false

	evalable := atoms.NewLeafValue(value)
	evaluated := evalable.Evaluate(context())

	if evaluated != value {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestTruthyConditionalValue(t *testing.T) {
	evalable := atoms.NewConditionalValue(
		atoms.NewLeafValue(1),
		atoms.NewLeafValue("a"),
		atoms.NewLeafValue("b"),
	)
	value := evalable.Evaluate(context())

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
	value := evalable.Evaluate(context())

	if value != "b" {
		t.Error("Evaluated value and expected value did not match")
	}

	// Test with the else branch being nil.
	evalable = atoms.NewConditionalValue(
		atoms.NewLeafValue(false),
		atoms.NewLeafValue("a"),
		nil,
	)
	value = evalable.Evaluate(context())

	if value != nil {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestNestedConditionalCreateEvalable(t *testing.T) {
	evalable := atoms.CreateEvalable(
		map[string]interface{}{
			"if": map[string]interface{}{
				"if":   true,
				"then": false,
				"else": true,
			},
			"then": "then-value",
			"else": "else-value",
		},
	)
	value := evalable.Evaluate(context())

	if value != "else-value" {
		t.Error("Evaluated value and expected value did not match")
	}

	// Same query as above, but negative case.
	evalable = atoms.CreateEvalable(
		map[string]interface{}{
			"if": map[string]interface{}{
				"if":   false,
				"then": false,
				"else": true,
			},
			"then": "then-value",
			"else": "else-value",
		},
	)
	value = evalable.Evaluate(context())

	if value != "then-value" {
		t.Error("Evaluated value and expected value did not match")
	}

	// Same query as above, but nil.
	evalable = atoms.CreateEvalable(
		map[string]interface{}{
			"if": map[string]interface{}{
				"if":   false,
				"then": false,
				"else": nil,
			},
			"then": "then-value",
			"else": "else-value",
		},
	)
	value = evalable.Evaluate(context())

	if value != "else-value" {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestCreateEvalableMissingBranch(t *testing.T) {
	// Missing if branch.
	evalable := atoms.CreateEvalable(
		map[string]interface{}{
			"if":   true,
			"else": "else-value",
		},
	)
	value := evalable.Evaluate(context())

	if value != nil {
		t.Error("Evaluated value and expected value did not match")
	}

	// Missing else branch.
	evalable = atoms.CreateEvalable(
		map[string]interface{}{
			"if":   false,
			"then": "then-value",
		},
	)
	value = evalable.Evaluate(context())

	if value != nil {
		t.Error("Evaluated value and expected value did not match")
	}

	// Missing both branches.
	evalable = atoms.CreateEvalable(
		map[string]interface{}{
			"if": false,
		},
	)
	value = evalable.Evaluate(context())

	if value != nil {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestStringConditional(t *testing.T) {
	// The conditional value is truthy.
	evalable := atoms.CreateEvalable(
		map[string]interface{}{
			"if":   "truthy",
			"then": "then-value",
			"else": "else-value",
		},
	)
	value := evalable.Evaluate(context())

	if value != "then-value" {
		t.Error("Evaluated value and expected value did not match")
	}

	// Conditional value is falsy.
	evalable = atoms.CreateEvalable(
		map[string]interface{}{
			"if":   "",
			"then": "then-value",
			"else": "else-value",
		},
	)
	value = evalable.Evaluate(context())

	if value != "else-value" {
		t.Error("Evaluated value and expected value did not match")
	}
}

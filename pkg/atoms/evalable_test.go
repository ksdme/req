package atoms_test

import (
	"testing"

	"github.com/ksdme/req/pkg/atoms"
)

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
	value := evalable.Evaluate(atoms.EmptyContext())

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
	value = evalable.Evaluate(atoms.EmptyContext())

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
	value = evalable.Evaluate(atoms.EmptyContext())

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
	value := evalable.Evaluate(atoms.EmptyContext())

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
	value = evalable.Evaluate(atoms.EmptyContext())

	if value != nil {
		t.Error("Evaluated value and expected value did not match")
	}

	// Missing both branches.
	evalable = atoms.CreateEvalable(
		map[string]interface{}{
			"if": false,
		},
	)
	value = evalable.Evaluate(atoms.EmptyContext())

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
	value := evalable.Evaluate(atoms.EmptyContext())

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
	value = evalable.Evaluate(atoms.EmptyContext())

	if value != "else-value" {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestConditionalWithVariables(t *testing.T) {
	evalable := atoms.CreateEvalable(
		map[string]interface{}{
			"if":   "<condition>",
			"then": "<then>",
			"else": "<else>",
		},
	)
	value := evalable.Evaluate(&atoms.Context{
		Variables: map[string]interface{}{
			"condition": false,
			"then":      "then-value",
			"else":      "else-value",
		},
	})

	if value != "else-value" {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestConditionalWithMissingVariables(t *testing.T) {
	evalable := atoms.CreateEvalable(
		map[string]interface{}{
			"if":   "<condition>",
			"then": "<then>",
			"else": "<else>",
		},
	)
	value := evalable.Evaluate(&atoms.Context{
		Variables: map[string]interface{}{
			"else": "else-value",
		},
	})

	if value != "else-value" {
		t.Error("Evaluated value and expected value did not match")
	}
}

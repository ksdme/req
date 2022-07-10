package atoms_test

import (
	"testing"

	"github.com/ksdme/req/pkg/atoms"
)

func TestStringLeafValue(t *testing.T) {
	var value interface{} = "string"

	evalable := atoms.NewLeafValue(value)
	evaluated := evalable.Evaluate(atoms.EmptyContext())

	if evaluated != value {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestIntLeafValue(t *testing.T) {
	var value interface{} = 1.23

	evalable := atoms.NewLeafValue(value)
	evaluated := evalable.Evaluate(atoms.EmptyContext())

	if evaluated != value {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestBoolLeafValue(t *testing.T) {
	var value interface{} = false

	evalable := atoms.NewLeafValue(value)
	evaluated := evalable.Evaluate(atoms.EmptyContext())

	if evaluated != value {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestSimpleVariableResolution(t *testing.T) {
	// Check if number resolution works.
	var value interface{} = 123

	context := &atoms.Context{
		Variables: map[string]interface{}{
			"value": value,
		},
	}
	evalable := atoms.NewLeafValue("<value>")
	evaluated := evalable.Evaluate(context)

	if evaluated != value {
		t.Error("Evaluated value and expected value did not match")
	}

	// Check if boolean variable resolution works.
	value = true

	context = &atoms.Context{
		Variables: map[string]interface{}{
			"value": value,
		},
	}
	evalable = atoms.NewLeafValue("<value>")
	evaluated = evalable.Evaluate(context)

	if evaluated != value {
		t.Error("Evaluated value and expected value did not match")
	}

	// Check if string resolution works.
	value = "hello-world"

	context = &atoms.Context{
		Variables: map[string]interface{}{
			"value": value,
		},
	}
	evalable = atoms.NewLeafValue("<value>")
	evaluated = evalable.Evaluate(context)

	if evaluated != value {
		t.Error("Evaluated value and expected value did not match")
	}
}

func TestStringInterpolationWorks(t *testing.T) {
	// String interpolation values.
	context := &atoms.Context{
		Variables: map[string]interface{}{
			"adjective": "big",
			"noun":      "bear",
		},
	}
	evalable := atoms.NewLeafValue("love <adjective> <noun>")
	evaluated := evalable.Evaluate(context)

	if evaluated != "love big bear" {
		t.Error("Evaluated value and expected value did not match")
	}

	// Check if non string data types interpolation works as expected too.
	context = &atoms.Context{
		Variables: map[string]interface{}{
			"number":  1,
			"boolean": true,
			"string":  "text",
		},
	}
	evalable = atoms.NewLeafValue("a <number>, <boolean>, <string>")
	evaluated = evalable.Evaluate(context)

	if evaluated != "a 1, true, text" {
		t.Errorf("Evaluated value and expected value did not match")
	}
}

func TestMissingVariable(t *testing.T) {
	context := &atoms.Context{
		Variables: map[string]interface{}{},
	}
	evalable := atoms.NewLeafValue("hey <name>")
	evaluated := evalable.Evaluate(context)

	// TODO: Should the resulting value be trimmed?
	if evaluated != "hey " {
		t.Errorf("Evaluated value and expected value did not match %v", evaluated)
	}
}

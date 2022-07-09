package atoms

import (
	"strings"
)

type Evalable interface {
	Evaluate(context *Context) interface{}
}

// Represents a terminal value. It evaluates to itself after interpolating
// it with the context.
type LeafValue struct {
	value interface{}
}

func NewLeafValue(value interface{}) *LeafValue {
	return &LeafValue{
		value: value,
	}
}

func (leaf *LeafValue) Evaluate(context *Context) interface{} {
	// TODO: Resolve variables from context here.
	return leaf.value
}

// Represents a conditional value. The value evaluates to the Then parameter
// if If value is truthy otherwise it returns an Else value.
type ConditionalValue struct {
	If   Evalable
	Then Evalable
	Else Evalable
}

func NewConditionalValue(ifEvalable Evalable, thenEvalable Evalable, elseEvalable Evalable) *ConditionalValue {
	return &ConditionalValue{
		If:   ifEvalable,
		Then: thenEvalable,
		Else: elseEvalable,
	}
}

func (conditional *ConditionalValue) Evaluate(context *Context) interface{} {
	var condition interface{} = nil
	if conditional.If != nil {
		condition = conditional.If.Evaluate(context)
	}

	var thenValue interface{} = nil
	if conditional.Then != nil {
		thenValue = conditional.Then.Evaluate(context)
	}

	var elseValue interface{} = nil
	if conditional.Else != nil {
		elseValue = conditional.Else.Evaluate(context)
	}

	if condition != nil {
		switch value := condition.(type) {
		case string:
			if strings.TrimSpace(value) != "" {
				return thenValue
			}

		case int, float32, float64:
			return thenValue

		case bool:
			if value {
				return thenValue
			}
		}
	}

	return elseValue
}

// Returns an evalable from a given piece of data. Automatically creates
// the nested structure when necessary.
func CreateEvalable(value interface{}) Evalable {
	// The terminal scalar values.
	switch typedValue := value.(type) {
	// Handle the case of nested objects.
	// TODO: What if the map keys are integers?
	case map[string]interface{}:
		// Support the conditional values.
		if ifValue, ok := typedValue["if"]; ok {
			thenValue, ok := typedValue["then"]
			if !ok {
				thenValue = nil
			}

			elseValue, ok := typedValue["else"]
			if !ok {
				elseValue = nil
			}

			return NewConditionalValue(
				CreateEvalable(ifValue),
				CreateEvalable(thenValue),
				CreateEvalable(elseValue),
			)
		}

	default:
		return NewLeafValue(value)
	}

	// TODO: Should throw an error instead?
	return NewLeafValue(nil)
}

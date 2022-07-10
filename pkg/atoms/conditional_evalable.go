package atoms

import "strings"

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

		case bool:
			if value {
				return thenValue
			}

		case int, float32, float64:
			return thenValue
		}
	}

	return elseValue
}

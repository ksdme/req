package atoms

type Evalable interface {
	Evaluate(context *Context) interface{}
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
		// TODO: Could we use something like map structure to deal with this instead?
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
	}

	return NewLeafValue(value)
}

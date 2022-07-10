package atoms

import (
	"fmt"
	"regexp"
	"strings"
)

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

// Evaluates the value with optional variable interpolation.
func (leaf *LeafValue) Evaluate(context *Context) interface{} {
	switch value := leaf.value.(type) {
	case string:
		value = strings.TrimSpace(value)

		// In case the value doesn't have any interpolation, the exact value from
		// the context will be returned.
		if variableTemplateSpan.MatchString(value) {
			return resolveVariable(value, context)
		} else {
			// In case the value has multiple variables, perform the interpolation.
			for _, match := range variableTemplate.FindAllString(value, -1) {
				resolution := resolveVariable(match, context)
				if resolution == nil {
					resolution = ""
				}

				replacement := fmt.Sprint(resolution)
				value = strings.ReplaceAll(value, match, replacement)
			}
		}

		return value
	}

	return leaf.value
}

// Regex template for the variable placeholder.
var variableTemplate = regexp.MustCompile(`<([\w-]+)>`)
var variableTemplateSpan = regexp.MustCompile(`^<([\w-]+)>$`)

// Given a placeholder, return its value.
func resolveVariable(placeholder string, context *Context) interface{} {
	variable := variableTemplate.ReplaceAllString(placeholder, "$1")

	value, ok := context.Variables[variable]
	if !ok {
		return nil
	}

	return value
}

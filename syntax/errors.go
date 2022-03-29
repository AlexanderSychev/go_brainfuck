package syntax

import "fmt"

// UnknownOperatorError returns when parser function receive unknown operator
type UnknownOperatorError struct {
	// Received operator
	value string
}

func (e UnknownOperatorError) Error() string {
	return fmt.Sprintf("Syntax error: invalid operator \"%s\"", e.value)
}

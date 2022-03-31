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

// UnexpectedTokenError returns when parser function receive operator with not closed or not opened cycle
type UnexpectedTokenError struct {
	// Received token
	value string
}

func (e UnexpectedTokenError) Error() string {
	return fmt.Sprintf("Syntax error: unexpected token \"%s\"", e.value)
}

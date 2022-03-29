package runtime

import (
	"fmt"
	"strings"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Corresponding command errors
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// SyntaxError returns when runtime cannot find command for received operator
type SyntaxError struct {
	operator string
}

func (err SyntaxError) Error() string {
	return fmt.Sprintf("Syntax error: unknown operator \"%s\"", err.operator)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// "Command" interface description and implementations
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Command is elementary execution unit for program (implementation of operator)
type Command interface {
	fmt.Stringer
	// Execute runs command body to received runtime memory
	Execute(memory *Memory) error
}

// Concrete commands

// ToNext is implementation of ">" operator - it set next cell as current
type ToNext struct{}

func (cmd ToNext) String() string {
	return ">"
}

func (cmd ToNext) Execute(memory *Memory) error {
	memory.ToNext()
	return nil
}

// ToPrevious is implementation of "<" operator - it set previous cell as current
type ToPrevious struct{}

func (cmd ToPrevious) String() string {
	return "<"
}

func (cmd ToPrevious) Execute(memory *Memory) error {
	memory.ToPrevious()
	return nil
}

// Increment is implementation of "+" operator - it increments value of current cell
type Increment struct{}

func (cmd Increment) String() string {
	return "+"
}

func (cmd Increment) Execute(memory *Memory) error {
	memory.SetValue(memory.GetValue() + 1)
	return nil
}

// Decrement is implementation of "-" operator - it decrements value of current cell
type Decrement struct{}

func (cmd Decrement) String() string {
	return "-"
}

func (cmd Decrement) Execute(memory *Memory) error {
	memory.SetValue(memory.GetValue() - 1)
	return nil
}

// Write is implementation of "." operator - it prints value of current cell
type Write struct{}

func (cmd Write) String() string {
	return "."
}

func (cmd Write) Execute(memory *Memory) error {
	sValue := string([]byte{memory.GetValue()})
	fmt.Print(sValue)
	return nil
}

// Read is implementation of "," operator - it reads value from user and stores it into current cell
type Read struct{}

func (cmd Read) String() string {
	return ","
}

func (cmd Read) Execute(memory *Memory) error {
	var err error = nil
	sValue := ""
	_, err = fmt.Scanln(&sValue)
	if err != nil {
		return err
	}

	chars := strings.Split(sValue, "")
	ascii := []byte(chars[0])
	if len(ascii) > 1 {
		return nil
	}

	memory.SetValue(ascii[0])
	return nil
}

// Cycle is implementation of "[" and "]" operators - it executes Cycle's body
// until value of current cell is bigger than 0
type Cycle []Command

func (cmd Cycle) String() string {
	items := make([]string, 0)
	for _, c := range cmd {
		items = append(items, c.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(items, ""))
}

func (cmd Cycle) Execute(memory *Memory) error {
	var err error = nil

	for memory.GetValue() > 0 {
		for _, c := range cmd {
			err = c.Execute(memory)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type Program []Command

func (cmd Program) String() string {
	items := make([]string, 0)
	for _, c := range cmd {
		items = append(items, c.String())
	}
	return strings.Join(items, "")
}

func (cmd Program) Execute(memory *Memory) error {
	var err error = nil

	for _, c := range cmd {
		err = c.Execute(memory)
		if err != nil {
			return err
		}
	}

	return nil
}

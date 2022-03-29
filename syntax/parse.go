package syntax

import (
	"go_brainfuck/runtime"
	"strings"
)

// canonize prepares source code to parsing - removes spaces, tabs and newlines
func canonize(src string) string {
	replacer := strings.NewReplacer("\n", "", "\r", "", "\t", "", " ", "")
	return replacer.Replace(src)
}

// explode transforms canonized source code to characters slice
func explode(src string) []string {
	return strings.Split(src, "")
}

// operatorsToCommands transforms
func operatorsToCommands(operators []string) ([]runtime.Command, error) {
	stack := NewStack()
	current := make([]runtime.Command, 0)

	for _, operator := range operators {
		switch operator {
		case ">":
			current = append(current, runtime.ToNext{})
		case "<":
			current = append(current, runtime.ToPrevious{})
		case "+":
			current = append(current, runtime.Increment{})
		case "-":
			current = append(current, runtime.Decrement{})
		case ",":
			current = append(current, runtime.Read{})
		case ".":
			current = append(current, runtime.Write{})
		case "[":
			stack.Push(current)
			current = make([]runtime.Command, 0)
		case "]":
			cmd := runtime.Cycle(current)
			current = stack.Pop()
			current = append(current, cmd)
		default:
			return current, UnknownOperatorError{operator}
		}
	}

	return current, nil
}

// Parse returns
func Parse(src string) (runtime.Program, error) {
	operators, err := operatorsToCommands(explode(canonize(src)))
	return operators, err
}

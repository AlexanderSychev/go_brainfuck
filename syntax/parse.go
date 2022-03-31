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

// checkCycles check that canonized source code does not contain unclosed cycles (single "[" or "]" operators)
func checkCycles(src string) error {
	var err error = nil

	cycleOpenCount := strings.Count(src, "[")
	cycleCloseCount := strings.Count(src, "]")
	if cycleOpenCount != cycleCloseCount {
		value := "["
		if cycleOpenCount < cycleCloseCount {
			value = "]"
		}
		err = UnexpectedTokenError{value}
	}

	return err
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
	canonized := canonize(src)

	err := checkCycles(canonized)
	if err != nil {
		return make([]runtime.Command, 0), err
	}

	operators, err := operatorsToCommands(explode(canonized))
	return operators, err
}

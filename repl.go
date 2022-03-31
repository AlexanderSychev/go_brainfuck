package main

import (
	"fmt"
	"go_brainfuck/runtime"
	"go_brainfuck/syntax"
	"strings"
)

const (
	commandExit = "exit"
	commandQuit = "quit"
	commandHelp = "help"
	commandHelpShort = "?"
)

func repl(memory *runtime.Memory) error {
	fmt.Printf("Go Brainfuck v%s REPL\n\n", VERSION)

	var err error = nil
	var program runtime.Program
	var cmd = ""

	for true {
		fmt.Print("> ")
		_, err = fmt.Scanln(&cmd)
		if err != nil {
			if strings.Contains(err.Error(), "unexpected newline") {
				continue
			} else {
				return err
			}
		}

		switch cmd {
		case commandHelp, commandHelpShort:
			fmt.Println("Usage:")
			fmt.Println("  - Type any valid Brainfuck code")
			fmt.Printf("  - Type \"%s\" or \"%s\" to exit\n", commandExit, commandQuit)
			fmt.Printf("  - Type \"%s\" or \"%s\" to get help\n", commandHelp, commandHelpShort)
		case commandExit, commandQuit:
			return nil
		default:
			program, err = syntax.Parse(cmd)
			if err != nil {
				return err
			}

			err = program.Execute(memory)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

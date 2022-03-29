package main

import (
	"fmt"
	"go_brainfuck/runtime"
	"go_brainfuck/syntax"
)

func RunREPL(memory *runtime.Memory) error {
	fmt.Printf("Go Brainfuck v1.0.0 REPL\n\nType \"exit\" to quit REPL\n\n")

	var err error = nil
	var program runtime.Program
	var cmd = ""

	for cmd != "exit" {
		fmt.Print("> ")
		_, err = fmt.Scanln(&cmd)
		if err != nil {
			return err
		}

		if cmd != "exit" {
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

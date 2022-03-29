package main

import (
	"github.com/urfave/cli"
	"go_brainfuck/runtime"
	"go_brainfuck/syntax"
	"log"
	"os"
)

const VERSION = "0.0.1"

func action(context *cli.Context) error {
	memory := runtime.NewMemory()

	if context.NArg() > 0 {
		filename := context.Args().First()

		raw, err := os.ReadFile(filename)
		if err != nil {
			return err
		}

		src := string(raw)
		program, err := syntax.Parse(src)
		if err != nil {
			return err
		}

		err = program.Execute(memory)
		if err != nil {
			return err
		}
	} else {
		err := RunREPL(memory)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "go_brainfuck"
	app.Action = action

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"github.com/urfave/cli"
	"go_brainfuck/runtime"
	"go_brainfuck/syntax"
	"os"
	"path"
)

// action is callback function for interpreter CLI
func action(context *cli.Context) error {
	memory := runtime.NewMemory()

	// If version flag provided then show version and stop application
	if context.Bool("version") {
		fmt.Printf("Go Brainfuck v%s", VERSION)
		return nil
	}

	if context.NArg() > 0 {
		workdir, err := os.Getwd()
		if err != nil {
			return err
		}

		filename := context.Args().First()

		raw, err := os.ReadFile(path.Join(workdir, filename))
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
		err := repl(memory)
		if err != nil {
			return err
		}
	}

	return nil
}

// newCLIApp is function which creates CLI application instance
func newCLIApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.BoolFlag{
			Name: "v, version",
			Usage: "Show version of interpreter",
		},
	}
	app.Name = "go_brainfuck"
	app.Action = action

	return app
}

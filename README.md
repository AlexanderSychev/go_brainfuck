# Go Brainfuck

Brainfuck interpreter written in Go.

* Has "infinite" memory (instead of 30000 bytes array);
* Can execute Brainfuck sourcecode files;
* Can be used as REPL;

## How to use

To run REPL:

```shell
go_brainfuck
```

To execute file with source code:

```shell
go_brainfuck <path-to-file>
```

## How to build

Clone repository to:

```shell
git clone git@github.com:AlexanderSychev/go_brainfuck.git
```

Select directory with cloned project:

```shell
cd <directory-with-go-brainfuck>
```

Build package

```shell
go build go_brainfuck
```

## Roadmap

* Do not interrupt REPL on errors;
* Add new-line adding for opened cycles;
* Byte code compilation;
* Dump "byte tape" to file (with position);
* Restore "byte tape" from file (with position);

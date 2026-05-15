# calc

> **Disclosure:** This repository — including its source code, tests, and this
> README — was built with the assistance of a large language model (LLM).
> Review the code before relying on it in production.

A small command-line calculator written in Go. It accepts an operator followed
by two or more numeric operands and prints the result to stdout.

## Dependencies

- Go 1.26.2 or newer (per `go.mod`).
- The Go standard library only — no third-party modules are imported.

The module declaration:

```
module calc
go 1.26.2
```

## Build

From the `calc/` directory:

```sh
go build -o calc .
```

This produces a `calc` binary in the current directory. To install into
`$GOBIN` (or `$GOPATH/bin`) instead:

```sh
go install .
```

## Usage

```
calc <operator> <num> <num> [num ...]
```

Supported operators:

| Operator | Aliases | Behavior                                          |
| -------- | ------- | ------------------------------------------------- |
| `add`    | `+`     | Sum of all operands.                              |
| `sub`    | `-`     | First operand minus each subsequent operand.      |
| `mul`    | `x`, `*`| Product of all operands.                          |
| `div`    | `/`     | First operand divided by each subsequent operand. |

Numbers are parsed as `float64`. At least two operands are required; fewer
operands, an unknown operator, a non-numeric argument, or a division by zero
exits with code `2` and prints a usage or error message to stderr.

### Examples

```sh
$ ./calc add 1 2
3

$ ./calc + 1 2 3 4
10

$ ./calc sub 10 1 2 3
4

$ ./calc - 5 2
3

$ ./calc add 0.1 0.2 0.3
0.6000000000000001

$ ./calc sub 1 5
-4

$ ./calc mul 2 3 4
24

$ ./calc x 1.5 2
3

$ ./calc div 100 2 5
10

$ ./calc / 7 2
3.5
```

### Exit codes

- `0` — success, result printed to stdout.
- `2` — usage error (too few args, unknown operator, unparseable number, or
  division by zero); diagnostic printed to stderr.

## Test

The test suite (`main_test.go`) builds the binary once via `TestMain` and then
exercises it as a black box through `os/exec`. From the `calc/` directory:

```sh
go test ./...
```

Useful variants:

```sh
go test -v ./...          # verbose, lists each subtest
go test -run TestCalc ./. # run only the TestCalc table
go test -race ./...       # enable the race detector
go test -cover ./...      # report coverage
```

The tests require `go` to be on `PATH` at test time because `TestMain` shells
out to `go build` to produce the binary under test in a temp directory.

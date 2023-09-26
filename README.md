# go-cat

a GNU cat alternative written in Go.

This is an experimental and educational project, not intended to fully replace the GNU `cat` command.

## Installation

```sh
go install github.com/jveillet/go-cat@latest
```

Go will automatically install it in your `$GOPATH/bin` directory which should be in your `$PATH`.

Once installed you should have the `go-cat` command available. Confirm by typing `go-cat` at a command line.

## Building from source

```sh
go build -o go-cat
```

Building without cgo (disables calling C code (import "C"))

```sh
CGO_ENABLED=0 go build -o go-cat
```

## Usage

```sh
$ go-cat -h
With no FILE, or when FILE is -, read standard input.

Usage:
  go-cat [OPTIONS]... [FILE]...

Flags:
  -h, --help     help for go-cat
  -n, --number   number all output lines
```

## Acknowledgements

This project uses [cobra](https://github.com/spf13/cobra/) to handle command line arguments.

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/antichris/go-pirev"
	"github.com/antichris/go-pirev/bits"
	"github.com/antichris/go-pirev/internal/print"
	"github.com/antichris/go-pirev/legacy"
)

const usage = `Usage:  %s REVCODE

Display details about a Raspberry Pi revision code.

  --help	display this help and exit

The REVCODE is a hexadecimal Raspberry Pi revision code. It can be either
new-style or old-style revision code and it can have any leading zeros omitted
(i.e. the REVCODE values "0013" and "13" both yield the same result)

Exit status:
 0  if OK,
 1  if REVCODE could not be parsed as a hexadecimal number,
 2  if old-style revision code hardware information could not be looked up.
`

// TODO Consider flags to control (filter) output of fields.

func main() { exit(trueMain(wrOut, wrErr, args)) }

var (
	args = os.Args
	exit = os.Exit

	wrOut io.Writer = os.Stdout
	wrErr io.Writer = os.Stderr
)

func trueMain(wr io.Writer, wrErr io.Writer, args []string) (exitCode int) {
	if len(args) < 2 || args[1] == "--help" {
		printUsage(wr, args)
		return
	}

	r, err := strconv.ParseUint(args[1], revBase, bits.CodeBitsize)
	if err != nil {
		printUsageWithError(wr, wrErr, args, err)
		return ExitParseErr
	}

	if i := pirev.Identify(pirev.Code(r)); i.NewStyleRev {
		if _, err := print.NewPrinter().Print(wr, i); err != nil {
			panic(err)
		}
		return
	}

	var i legacy.Info
	if i, err = legacy.Describe(legacy.RevCode(r)); err != nil {
		printUsageWithError(wr, wrErr, args, err)
		return ExitLookupErr
	}

	if _, err = print.NewLegacyPrinter().Print(wr, i); err != nil {
		panic(err)
	}
	return
}

func printUsage(wr io.Writer, args []string) {
	if _, err := fmt.Fprintf(wr, usage, args[0]); err != nil {
		panic(err)
	}
}

const revBase = 16 // The number base used in revision codes.

func printUsageWithError(wr io.Writer, wrErr io.Writer, args []string, err error) {
	if _, err := fmt.Fprintln(wrErr, err); err != nil {
		panic(err)
	}
	printUsage(wr, args)
}

const (
	// ExitOK is the default exit code and means that no errors occurred.
	ExitOK = iota
	// ExitParseErr means arg couldn't be parsed as a hexadecimal number.
	ExitParseErr
	// ExitLookupErr means old-style revision info couldn't be looked up.
	ExitLookupErr
)

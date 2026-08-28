package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/jghiloni/strip-unicode/unistripper"
)

var version = "0.0.0-dev"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var versionFlag bool
	flag.CommandLine.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Usage: strip-unicode < infile > outfile") //nolint:errcheck
		flag.CommandLine.PrintDefaults()
	}
	flag.CommandLine.BoolVar(&versionFlag, "version", false, "Show the version and exit")
	flag.Parse()

	if versionFlag {
		fmt.Fprintf(flag.CommandLine.Output(), "Version %s\n\n", version) //nolint:errcheck
		os.Exit(0)
	}

	if err := unistripper.StripUnicode(ctx, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error stripping unicode: %v", err)
		os.Exit(1)
	}
}

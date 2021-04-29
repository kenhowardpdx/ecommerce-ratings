package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kenhowardpdx/ecommerce-ratings/internal/server"
)

var version = os.Getenv("VERSION")

func main() {
	if err := Run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// Run is an abstraction for main that enables testing.
func Run(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("ecommerce-ratings", flag.ContinueOnError)
	flags.SetOutput(stderr)
	useLocalhost := flags.Bool("localhost", false, "sets server address to localhost")
	port := flags.Int("port", 8080, "http port")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	serverAddress := "0.0.0.0"
	if *useLocalhost {
		serverAddress = "localhost"
	}
	if version == "" {
		version = "0.0.0"
	}
	fmt.Printf("starting server at %s:%d\n", serverAddress, *port)
	srv := server.Server{
		Address: serverAddress,
		Port:    *port,
		Version: version,
	}
	srv.Start()
	return nil
}

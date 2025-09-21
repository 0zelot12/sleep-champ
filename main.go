package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'start', 'status' or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "start":
		startCmd := flag.NewFlagSet("start", flag.ExitOnError)
		delay := startCmd.Int("delay", 60, "Delay in seconds before shutdown")

		startCmd.Parse(os.Args[2:])

		fmt.Println("Starting with delay:", *delay)

	case "status":
		fmt.Println("Status: running")

	case "delete":
		fmt.Println("Shutdown canceled")

	default:
		fmt.Println("expected 'start', 'status' or 'delete' subcommands")
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	todayCmd := flag.NewFlagSet("today", flag.ExitOnError)
	yesterdayCmd := flag.NewFlagSet("yesterday", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected you to specify either 'today' or 'yesterday")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "today":
		todayCmd.Parse(os.Args[2:])
		// need to call my request here
		fmt.Println("placeholder for today")
	case "yesterday":
		yesterdayCmd.Parse(os.Args[2:])
		// need to call my request here
		fmt.Println("placeholder for yesterday")
	default:
		fmt.Fprintln(os.Stderr, "Invalid operation")
		os.Exit(1)
	}
}
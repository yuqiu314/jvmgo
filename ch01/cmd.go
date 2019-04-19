package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "Help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "class parth")
	flag.StringVar(&cmd.cpOption, "cp", "", "class path")
	flag.Parse()
	args := flag.Args()
	if len(args) >0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
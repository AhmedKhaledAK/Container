package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)


// docker 				run image <cmd> <params>
// go run container.go 	run 	  <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])
}
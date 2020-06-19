package main

import (
	"os"
	"testing"
)

func Test(t *testing.T){
	args := os.Args[0:1] // Name of the program.
	args = append(args, "") // Append a flag
	run(args)
 }
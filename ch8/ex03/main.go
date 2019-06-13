package main

import (
	"os"
	"syscall"
)

func main() {
	p, err := os.FindProcess(8964)
	if err != nil {
		panic(err) }
	if err = p.Signal(syscall.SIGTERM); err != nil {
		panic(err)
	}
}
//go:build linux

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		fmt.Println("default")
	}
}

func run() {
	fmt.Println("Running...")
	fmt.Printf("arguments: %v", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	go func() {
		io.Copy(os.Stdout, stdout)
		io.Copy(os.Stderr, stderr)
	}()

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Run()
}

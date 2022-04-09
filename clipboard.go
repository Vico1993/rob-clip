package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func GetValue() string {
	cmd := exec.Command("pbpaste")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
    if err != nil {
		log.Fatal(err)
        os.Exit(1)
	}

	buf := bufio.NewReader(stdout) // Notice that this is not in a loop

	line, _, _ := buf.ReadLine()

	cmd.Process.Kill()

	return string(line)
}
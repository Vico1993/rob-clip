package main

import "github.com/Vico1993/rob-clip/cmd"

var (
  list = []Copyed{}
)

func main() {
	startDaemon()

	cmd.Execute()
}

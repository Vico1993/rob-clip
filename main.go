package main

import "github.com/Vico1993/rob-clip/cmd"

var (
  list = []string{}
)


func main() {
	// Initialisation of the config
	initConfig()

	startDaemon()

	cmd.Execute(list)
}

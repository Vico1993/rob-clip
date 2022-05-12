package main

import (
	"github.com/Vico1993/rob-clip/cmd"
	"github.com/Vico1993/rob-clip/config"
	"github.com/Vico1993/rob-clip/daemon"
)


func main() {
	// Initialisation of the config
	config.InitConfig()

	daemon.StartDaemon()

	cmd.Execute()
}

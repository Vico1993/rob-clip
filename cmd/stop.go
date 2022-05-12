package cmd

import (
	"log"
	"os"
	"syscall"

	"github.com/Vico1993/rob-clip/config"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop deamon",
	Run: func(cmd *cobra.Command, args []string) {
		pidFileName := config.GetConfigFolder() + "/daemon.pid"
		pid, err := daemon.ReadPidFile(pidFileName);
		if err != nil {
			log.Println("Error stopping daemon", err.Error())
		}

		err = syscall.Kill(pid, syscall.SIGKILL)
		if err != nil {
			log.Println("Error stopping daemon", err.Error())
		}

		e := os.Remove(pidFileName)
		if e != nil {
			log.Println("Error deleting daemon.pid", err.Error())
		}

		config.UpdateDaemonStatus(false)
		cmd.Print("Daemon killed")
	},
}
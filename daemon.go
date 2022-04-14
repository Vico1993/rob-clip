package main

import (
	"log"
	"os"
	"syscall"
	"time"

	"github.com/sevlyar/go-daemon"
	"github.com/spf13/viper"
)

var (
	start = false
	stop = make(chan struct{})
	done = make(chan struct{})
)

func findWordInList(word string) bool {
    for i := 0; i < len(list); i++ {
        if list[i] == word {
            return false
        }
    }

    return true
}

func worker() {
LOOP:
	for {
		log.Println("Populating model...")

		word := GetValue()

		if (findWordInList(word)) {
			list = append(list, GetValue())
		}

		// Every Second
		time.Sleep(time.Second)

		select {
		case <-stop:
			break LOOP
		default:
		}
	}
	done <- struct{}{}
}

func startDaemon() {
	cntxt := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[rob-clip]"},
	}

	deamon := viper.GetBool("DEAMON_STARTED")
	log.Println(deamon)

	if deamon {
		log.Println("Yoohoo ici")
		d, err := cntxt.Search()
		if err != nil {
			log.Printf("Unable send signal to the daemon: %s", err.Error())
		}

		daemon.SendCommands(d)
		return
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatalln(err)
	}
	if d != nil {
		return
	}

	defer cntxt.Release()

	log.Println("- - - - - - - - - - - - - - -")
	log.Println("daemon started")

	viper.Set("DEAMON_STARTED", "true")
	err = viper.WriteConfig()

	if err != nil {
		log.Fatalln(err.Error())
	}

	go worker()

	err = daemon.ServeSignals()
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	log.Println("daemon terminated")

	log.Println("DEAMONE WAS REBORN")
}

func stopDaemon(sig os.Signal) error {
	log.Println("terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}
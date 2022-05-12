package daemon

import (
	"log"
	"time"

	"github.com/Vico1993/rob-clip/config"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/viper"
)

var (
	stop = false
	done = make(chan struct{})
	list = []string{}
)

func isWordAlreadyInList(word string) bool {
    for i := 0; i < len(list); i++ {
        if list[i] == word {
            return true
        }
    }

    return false
}

func worker() {
LOOP:
	for {
		word := GetValue()

		if (!isWordAlreadyInList(word)) {
			list = append(list, GetValue())

			viper.Set("DAEMON_WORD", list)
			err := viper.WriteConfigAs(config.GetConfigFilePath())
			if err != nil {
				log.Fatal("Can't write value in config file at " + err.Error())
			}
		}

		// Every Second
		time.Sleep(time.Second)

		if stop {
			break LOOP
		}
	}
	done <- struct{}{}
}

func StartDaemon() {
	cntxt := &daemon.Context{
		PidFileName: config.GetConfigFolder() + "/daemon.pid",
		PidFilePerm: 0644,
		LogFileName: config.GetConfigFolder() + "/daemon.log",
		LogFilePerm: 0640,
		WorkDir:     config.GetConfigFolder(),
		Umask:       027,
		Args:        []string{"[rob-clip]"},
	}

	if viper.GetBool("DEAMON_STARTED") {
		d, err := cntxt.Search()
		if err != nil {
			log.Printf("Unable send signal to the daemon: %s", err.Error())
		}

		err = daemon.SendCommands(d)
		if err != nil {
			log.Fatalln("Can't send commands: " + err.Error())
		}
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
}
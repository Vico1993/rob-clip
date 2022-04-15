package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func getConfigFileName() (string) {
	return "/config.json"
}

func getConfigFilePath() (string) {
	return getConfigFolder() + getConfigFileName()
}

func getConfigFolder() (string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		homedir = "./"
	}

	return homedir + "/.rob-clip"
}

func initConfig () {
	path := getConfigFolder()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal("Can't create Folder at " + path)
		}
	}

	configFilePath := getConfigFilePath()

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		var file, err = os.Create(configFilePath)
        if err != nil {
            log.Fatal("Can't create config file at " + configFilePath)
        }
        defer file.Close()

		// Set default
		viper.Set("DEAMON_STARTED", false)
		viper.WriteConfigAs(configFilePath)
	}

	viper.SetConfigFile(getConfigFilePath())

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: %w \n", err)
	}
}
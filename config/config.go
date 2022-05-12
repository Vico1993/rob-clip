package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetConfigFileName() (string) {
	return "/config.json"
}

func GetConfigFilePath() (string) {
	return GetConfigFolder() + GetConfigFileName()
}

func GetConfigFolder() (string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		homedir = "./"
	}

	return homedir + "/.rob-clip"
}

func InitConfig () {
	path := GetConfigFolder()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal("Can't create Folder at " + path)
		}
	}

	configFilePath := GetConfigFilePath()

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		var file, err = os.Create(configFilePath)
        if err != nil {
            log.Fatal("Can't create config file at " + configFilePath)
        }
        defer file.Close()

		// Set default
		UpdateDaemonStatus(false)
	}

	viper.SetConfigFile(GetConfigFilePath())

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: %w \n", err)
	}
}

func UpdateDaemonStatus(status bool) {
	// Set default
	viper.Set("DEAMON_STARTED", status)
	err := viper.WriteConfigAs(GetConfigFilePath())
	if err != nil {
		log.Fatal("Can't write value in config file at " + err.Error())
	}
}
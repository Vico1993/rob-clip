package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
  list = []string{}
)

func initConfig () {
	path := "$HOME/.rob-clip"

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetDefault("DEAMON_STARTED", false)

	la, err := os.Stat(path);
	if err != nil {
		fmt.Println( "Error 2", os.IsNotExist(err))

		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0755)

			if err != nil {
				fmt.Println("Error 3", err.Error())
			}
		}
		os.Exit(1)
	}

	fmt.Println(la)

	// if _, err := os.Stat(path); os.IsNotExist(err) {
	// 	err = os.MkdirAll(path, 0755)

	// 	fmt.Println("AIE")

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	fmt.Println("here")
}

func main() {
	initConfig()

	// startDaemon()

	// cmd.Execute(list)
}

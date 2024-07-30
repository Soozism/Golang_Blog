package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal("config file dose not exist in ./config/config")
		} else {
			// Config file was found but another error was produced
			log.Fatal("Error reading the config :", err)
		}
		os.Exit(1)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

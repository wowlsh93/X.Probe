/*
Copyright Monitoring Corp. All Rights Reserved.

Written by HAMA
*/

package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func InitConfig(configPath string) (Configuration, error) {

	var configuration Configuration

	viper.SetConfigType("yaml")
	viper.SetConfigName("config") // name of config file (without extension)

	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		home, _ := homedir.Dir()
		configPath = home + "/monitoring/configurations"
		viper.AddConfigPath(configPath)
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
		return configuration, err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return configuration, err
	}

	err := viper.Unmarshal(&configuration)

	if err != nil {
		fmt.Printf("unable to decode into configuration struct, %s", err)
		return configuration, err
	}

	return configuration, nil
}

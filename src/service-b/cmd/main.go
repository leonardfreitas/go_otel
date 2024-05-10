/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"leonardfreitas/go_otel/src/service-b/cmd/cmd"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	cmd.Execute()
}

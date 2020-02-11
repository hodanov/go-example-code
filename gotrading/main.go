package main

import (
	"app/gotrading/bitflyer"
	"app/gotrading/config"
	"app/gotrading/utils"
	"fmt"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.APIKey, config.Config.APISecret)
	fmt.Println(apiClient.GetBalance())
}

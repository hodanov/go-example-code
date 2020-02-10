package main

import (
	"app/gotrading/config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.APIKey)
	fmt.Println(config.Config.APISecret)
}

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Test if viper is able to parse map structure in yml file in order, the result is NO
// go run test.go
func main() {
	fmt.Println("test viper get map")
	viper.SetConfigFile(".ayi.yml")
	viper.ReadInConfig()
	fmt.Println(viper.Get("debug"))
	fmt.Println(viper.Get("git.hosts")) // NOTE: The order is random
	// @gaocegege said it's possible to return slice instead of map, so order is kept.
}

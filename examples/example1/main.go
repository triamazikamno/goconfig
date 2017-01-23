/*
Example without configuration file. Only parameters and environment variables.
*/
package main

import (
	"fmt"

	"github.com/crgimenes/goConfig"
)

type mongoDB struct {
	Host string `cfgDefault:"example.com"`
	Port int    `cfgDefault:"999"`
}

type systemUser struct {
	Name     string `cfg:"name"`
	Password string `cfg:"passwd"`
}

type configTest struct {
	Domain  string
	User    systemUser `cfg:"user"`
	MongoDB mongoDB
}

func main() {
	config := configTest{}

	goConfig.PrefixEnv = "EXAMPLE"
	err := goConfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n\n%#v\n\n", config)

}

package main

import "fmt"
import "github.com/crgimenes/goConfig"
import _ "github.com/crgimenes/goConfig/json"

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
	fmt.Println("init")
	config := configTest{}

	goConfig.File = "config.json"
	goConfig.PrefixEnv = "EXAMPLE"
	err := goConfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n\n%#v\n\n", config)

	fmt.Println("end")
}

package main

import "fmt"
import "github.com/crgimenes/goConfig"

type mongoDB struct {
	Host string `cfgDefault:"example.com"`
	Port int    `cfgDefault:"999"`
}

type SystemUser struct {
	Name     string `cfg:"name"`
	Password string `cfg:"passwd"`
}

type configTest struct {
	Domain  string
	User    SystemUser `cfg:"user"`
	MongoDB mongoDB
}

func main() {
	fmt.Println("init")
	config := configTest{}

	err := goConfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n\n%#v\n\n", config)

	fmt.Println("end")
}

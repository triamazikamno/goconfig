/*
Example with configuration file.
*/
package main

import (
	"encoding/json"

	"github.com/crgimenes/goConfig"
	_ "github.com/crgimenes/goConfig/toml"
)

type mongoDB struct {
	Host string `toml:"host" cfg:"host" cfgDefault:"example.com"`
	Port int    `toml:"port" cfg:"port" cfgDefault:"999"`
}

type systemUser struct {
	Name     string `toml:"name" cfg:"name"`
	Password string `toml:"passwd" cfg:"passwd"`
}

type configTest struct {
	DebugMode bool       `toml:"debug" cfg:"debug" cfgDefault:"false"`
	Domain    string     `toml:"domain"`
	User      systemUser `toml:"user" cfg:"user"`
	MongoDB   mongoDB    `toml:"mongodb" cfg:"mongodb"`
}

func main() {
	config := configTest{}

	goConfig.File = "config.toml"
	err := goConfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	// just print struct on screen
	j, _ := json.MarshalIndent(config, "", "\t")
	println(string(j))
}

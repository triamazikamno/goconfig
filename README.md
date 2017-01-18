# goConfig
[![Go Report Card](https://goreportcard.com/badge/github.com/crgimenes/goConfig)](https://goreportcard.com/report/github.com/crgimenes/goConfig)

Load settings form parameters, environment variables and a configuration file.

## Example

```go
package main

import "fmt"
import "github.com/crgimenes/goConfig"

type mongoDB struct {
	Host string `cfgDefault:"example.com"`
	Port int    `cfgDefault:"999"`
}

type configTest struct {
	Domain  string
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
```
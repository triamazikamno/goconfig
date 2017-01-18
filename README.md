# goConfig
[![Go Report Card](https://goreportcard.com/badge/github.com/crgimenes/goConfig)](https://goreportcard.com/report/github.com/crgimenes/goConfig)
[![codecov](https://codecov.io/gh/crgimenes/goConfig/branch/master/graph/badge.svg)](https://codecov.io/gh/crgimenes/goConfig)
[![GoDoc](https://godoc.org/github.com/crgimenes/goConfig?status.png)](https://godoc.org/github.com/crgimenes/goConfig)

Load settings form parameters, environment variables and a configuration file.

## Install

```
go get github.com/crgimenes/goConfig
```

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
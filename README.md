# goConfig
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b8c6bf828cfa451d907321aded911509)](https://www.codacy.com/app/crgimenes/goConfig?utm_source=github.com&utm_medium=referral&utm_content=crgimenes/goConfig&utm_campaign=badger)
[![Build Status](https://travis-ci.org/crgimenes/goConfig.svg?branch=master)](https://travis-ci.org/crgimenes/goConfig)
[![Go Report Card](https://goreportcard.com/badge/github.com/crgimenes/goConfig)](https://goreportcard.com/report/github.com/crgimenes/goConfig)
[![codecov](https://codecov.io/gh/crgimenes/goConfig/branch/master/graph/badge.svg)](https://codecov.io/gh/crgimenes/goConfig)
[![GoDoc](https://godoc.org/github.com/crgimenes/goConfig?status.png)](https://godoc.org/github.com/crgimenes/goConfig)
[![Go project version](https://badge.fury.io/go/github.com%2Fcrgimenes%2FgoConfig.svg)](https://badge.fury.io/go/github.com%2Fcrgimenes%2FgoConfig)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)


goConfig uses a struct as input and populates the fields of this struct with parameters fom command line, environment variables and configuration file.

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
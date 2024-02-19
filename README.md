# tipper

[![CI](https://github.com/winebarrel/tipper/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/tipper/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/winebarrel/tipper.svg)](https://pkg.go.dev/github.com/winebarrel/tipper)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/winebarrel/tipper)](https://github.com/winebarrel/tipper/tags)
[![Go Report Card](https://goreportcard.com/badge/github.com/winebarrel/tipper)](https://goreportcard.com/report/github.com/winebarrel/tipper)

A library that recursively dumps structure information.

## Installation

```sh
go get github.com/winebarrel/tipper
```

### Usage

```go
package main

import (
	"fmt"

	"github.com/winebarrel/tipper"
)

type config struct {
	Home string     `env:"HOME"`
	Port int        `env:"PORT" envDefault:"3000"`
	Bar  *subconfig `envPrefix:"SUB_"`
}

type subconfig struct {
	Password     string `env:"PASSWORD,unset"`
	IsProduction bool   `env:"PRODUCTION"`
}

func main() {
	var c config
	ss := tipper.Dump(c)
	fmt.Println(ss[0].Fields[0]) //=> "{Password string [{env PASSWORD [unset]}]}"
	fmt.Println(ss)
}
```

```json
[
  {
    "name": "main.subconfig",
    "fields": [
      {
        "name": "Password",
        "type": "string",
        "tags": [
          {
            "key": "env",
            "name": "PASSWORD",
            "options": [
              "unset"
            ]
          }
        ]
      },
      {
        "name": "IsProduction",
        "type": "bool",
        "tags": [
          {
            "key": "env",
            "name": "PRODUCTION",
            "options": null
          }
        ]
      }
    ]
  },
  {
    "name": "main.config",
    "fields": [
      {
        "name": "Home",
        "type": "string",
        "tags": [
          {
            "key": "env",
            "name": "HOME",
            "options": null
          }
        ]
      },
      {
        "name": "Port",
        "type": "int",
        "tags": [
          {
            "key": "env",
            "name": "PORT",
            "options": null
          },
          {
            "key": "envDefault",
            "name": "3000",
            "options": null
          }
        ]
      },
      {
        "name": "Bar",
        "type": "*main.subconfig",
        "tags": [
          {
            "key": "envPrefix",
            "name": "SUB_",
            "options": null
          }
        ]
      }
    ]
  }
]
```

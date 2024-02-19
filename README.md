# tipper

[![CI](https://github.com/winebarrel/tipper/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/tipper/actions/workflows/ci.yml)

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

package main

import (
	"fmt"

	"github.com/winebarrel/tipper"
)

type config struct {
	Home string     `env:"HOME, required"`
	Port int        `env:"PORT" envDefault:"3000"`
	Bar  *subconfig `envPrefix:"SUB_"`
}

type subconfig struct {
	Password     string `env:"PASSWORD,unset,required"`
	IsProduction bool   `env:"PRODUCTION"`
}

func main() {
	var c config
	ss := tipper.Dump(c)
	fmt.Println(ss[0].Fields[0]) //=> "{Password string [{env PASSWORD [unset required]}]}"
	fmt.Println(ss)

	fromT := tipper.DumpT[config]()
	fmt.Println(fromT) // Same output as "fmt.Println(ss)"
}

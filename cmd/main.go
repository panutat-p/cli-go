package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/urfave/cli/v2"

	"cli-go/internal"
)

func main() {
	var conf internal.Config
	err := env.Parse(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("conf: %+v\n", conf)

	h := internal.NewHandler(&conf)

	app := &cli.App{
		Name:   "hello",
		Usage:  "say hello",
		Action: h.Hello,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "hello",
		Usage:  "say hello",
		Action: Hello,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func Hello(c *cli.Context) error {
	fmt.Println("hello")
	fmt.Println("arg 0:", c.Args().Get(0))
	fmt.Println("arg 1:", c.Args().Get(1))
	return nil
}

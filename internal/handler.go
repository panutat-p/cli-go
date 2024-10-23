package internal

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type Handler struct {
	Config *Config
}

func NewHandler(config *Config) *Handler {
	return &Handler{
		Config: config,
	}
}

func (h *Handler) GetConfig() *Config {
	return h.Config
}

func (h *Handler) Hello(c *cli.Context) error {
	fmt.Println("hello")
	fmt.Println("arg 0:", c.Args().Get(0))
	fmt.Println("arg 1:", c.Args().Get(1))
	return nil
}

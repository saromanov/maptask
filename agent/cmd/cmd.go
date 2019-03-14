package main

import (
	"context"
	"fmt"
	"os"
	"errors"

	"github.com/saromanov/maptask/agent/master"
	"github.com/saromanov/maptask/agent"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

// Config defines configuration for maptask
type Config struct {
	MasterServer string
}

// parseConfig provides parsing of the config .yml file
func parseConfig(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	var c *Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .config.yml: %v", err)
	}

	return c, nil
}

func run(c *Config) {
	if c.MasterServer != "" {
		master.Run(c.MasterServer)
	}

	ag := agent.New(c)
	if err := ag.Start(); err != nil {
		panic(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "maptask"
	app.Usage = "Distributed map reduce"
	app.Commands = []cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "path to .yml config",
			Action: func(c *cli.Context) error {
				configPath := c.Args().First()
				config, err := parseConfig(configPath)
				if err != nil {
					panic(err)
				}
				if err := run(config); err != nil {
					panic(err)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoTemplate-Plugin"
	app.Usage = "A plugin for go-template that accepts specific parameters."
	
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "template, t",
			Usage:  "Path to the template file.",
			EnvVar: "PLUGIN_TEMPLATE",
		},
		cli.StringFlag{
			Name:   "values, f",
			Usage:  "Path to the values file.",
			EnvVar: "PLUGIN_VALUES",
		},
		cli.StringFlag{
			Name:   "output, o",
			Usage:  "Path to the output directory.",
			EnvVar: "PLUGIN_OUTPUT",
		},
	}

	app.Action = func(c *cli.Context) error {
		runPlugin(
			c.String("template"),
			c.String("values"),
			c.String("output"),
		)
		return nil
	}

	app.Run(os.Args)
}

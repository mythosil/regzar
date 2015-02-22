package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "regzar"
	app.Usage = "Remote Control for REGZA"
	app.Version = "0.1"
	app.Author = "mythosil"
	app.Email = "mythosil@yahoo.co.jp"
	app.Commands = Commands
	return app
}

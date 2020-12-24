package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		UseShortOptionHandling: true,
		Flags:                  []cli.Flag{},
		Commands: []*cli.Command{
			Set,
			RemoveCmd,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

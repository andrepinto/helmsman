package main

import (
	"os"
	"fmt"
	app2 "github.com/andrepinto/helmsman/cmd/server/app"
)

func main() {
	app := app2.NewCliApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
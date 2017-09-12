package app

import (
	"gopkg.in/urfave/cli.v1"
)

const  REPO_URL = "HELMSMAN_REPO_URL"

type WChestCmdOptions struct {
	Port int
	RepoDir string
	RepoUrl string
	Debug bool
}

func NewWChestCmdOptions() *WChestCmdOptions{
	return &WChestCmdOptions{
	}
}

func (opts *WChestCmdOptions) AddFlags(app *cli.App){

	flags := []cli.Flag{
		cli.IntFlag{
			Name:        "port",
			Value:       8000,
			Usage:       "server port",
			Destination: &opts.Port,
		},
		cli.StringFlag{
			Name:        "repo.dir",
			Value:       "./charts",
			Usage:       "helm repo dir",
			Destination: &opts.RepoDir,
		},
		cli.StringFlag{
			Name:        "repo.url",
			Value:       "localhost:8000/charts/",
			Usage:       "helm repo url",
			EnvVar: 	 REPO_URL,
			Destination: &opts.RepoUrl,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Debug mode default to false",
			Destination: &opts.Debug,
		},

	}

	app.Flags = append(app.Flags, flags...)
}
package app

import (
	"gopkg.in/urfave/cli.v1"
)

const repoURL = "HELMSMAN_REPO_URL"

//HemlCmdOptions ...
type HemlCmdOptions struct {
	Port    int
	RepoDir string
	RepoUrl string
	Debug   bool
	Envs    []string
}

//NewHemlCmdOptionsCmdOptions ...
func NewHemlCmdOptionsCmdOptions() *HemlCmdOptions {
	return &HemlCmdOptions{}
}

//AddFlags ...
func (opts *HemlCmdOptions) AddFlags(app *cli.App) {

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
			EnvVar:      repoURL,
			Destination: &opts.RepoUrl,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Debug mode default to false",
			Destination: &opts.Debug,
		},
		cli.StringSliceFlag{
			Name:  "env",
			Value: &cli.StringSlice{"stable"},
		},
	}

	app.Flags = append(app.Flags, flags...)
}

package app

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/andrepinto/helmsman/api"
	"github.com/andrepinto/helmsman/pkg"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
)

//NewCliApp ...
func NewCliApp() *cli.App {

	app := cli.NewApp()

	app.Name = "helmsman"
	app.Version = VERSION

	opts := NewHemlCmdOptionsCmdOptions()
	opts.AddFlags(app)

	app.Action = func(c *cli.Context) error {

		if opts.Debug {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}

		opts.Envs = c.StringSlice("env")

		err := Init(opts)
		if err != nil {
			log.Fatal(err)
		}

		proc := api.NewServer(&api.ServerOptions{
			Port:    opts.Port,
			RepoDir: opts.RepoDir,
			RepoUrl: opts.RepoUrl,
		})
		log.Debug(opts)
		return proc.Run()
	}

	// sort flags by name
	sort.Sort(cli.FlagsByName(app.Flags))

	return app
}

//Init ...
func Init(opts *HemlCmdOptions) error {

	log.Info(opts.Envs)

	for _, env := range opts.Envs {
		folder := filepath.Join(opts.RepoDir, env)
		err := os.MkdirAll(folder, 0777)
		if err != nil {
			return err
		}

		err = pkg.Index(folder, folder, "")
		if err != nil {
			return err
		}

	}

	return nil

}

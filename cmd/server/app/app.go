package app

import (
	"sort"
	"gopkg.in/urfave/cli.v1"
	"github.com/andrepinto/helmsman/api"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/andrepinto/helmsman/pkg"
)

func NewCliApp() *cli.App {

	app := cli.NewApp()

	app.Name = "helmsman"
	app.Version = VERSION

	opts := NewWChestCmdOptions()
	opts.AddFlags(app)


	app.Action = func(c *cli.Context) error {

		if opts.Debug{
			log.SetLevel(log.DebugLevel)
		}else{
			log.SetLevel(log.InfoLevel)
		}

		err := Init(opts)
		if err != nil{
			log.Fatal(err)
		}

		proc := api.NewApiServer(&api.ApiServerOptions{
			Port: opts.Port,
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


func Init(opts *WChestCmdOptions) error{
	err := os.MkdirAll(opts.RepoDir, 0777)
	if err != nil {
		return err
	}

	err = pkg.Index(opts.RepoDir, opts.RepoDir, "")

	return err
}
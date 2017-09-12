package pkg

import (
	"path/filepath"
	"fmt"
	repo "k8s.io/helm/pkg/repo"
	"k8s.io/helm/pkg/chartutil"
	"os"
	"github.com/sirupsen/logrus"
)

func Index(dir, url, mergeTo string) error {
	out := filepath.Join(dir, "index.yaml")
	logrus.Debug(out)
	i, err := repo.IndexDirectory(dir, url)
	if err != nil {
		return err
	}
	if mergeTo != "" {
		i2, err := repo.LoadIndexFile(mergeTo)
		if err != nil {
			return fmt.Errorf("Merge failed: %s", err)
		}
		i.Merge(i2)
	}
	i.SortEntries()
	return i.WriteFile(out, 0755)
}

func Package(path string, destination string) error{
	path, err := filepath.Abs(path)

	ch, err := chartutil.LoadDir(path)
	if err != nil {
		return err
	}

	if filepath.Base(path) != ch.Metadata.Name {
		return fmt.Errorf("directory name (%s) and Chart.yaml name (%s) must match", filepath.Base(path), ch.Metadata.Name)
	}

	_, err = chartutil.LoadRequirements(ch)

	var dest string
	if destination == "." {
		// Save to the current working directory.
		dest, err = os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// Otherwise save to set destination
		dest = destination
	}

	_, err = chartutil.Save(ch, dest)

	return err
}
package repo

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/andrepinto/helmsman/pkg"
	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
	"k8s.io/helm/pkg/urlutil"
)

func (pr *RepoResource) chartCtrl(request *restful.Request, response *restful.Response) {

	log.Debug("get chart")

	id := request.PathParameter("chart")
	env := request.PathParameter("env")

	response.AddHeader("Content-Type", "text/plain")

	file, err := os.Open(filepath.Join(pr.RepoDir, env, id))
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, "500: Charts error.")
		return
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, "500: Charts error.")
		return
	}

	response.Write(b)
}

func (pr *RepoResource) uploadChartCtrl(request *restful.Request, response *restful.Response) {

	log.Debug("upload a chart")

	id := request.PathParameter("chart")
	env := request.PathParameter("env")

	response.AddHeader("Content-Type", "text/plain")

	f, err := os.Create(filepath.Join(pr.RepoDir, env, id))
	defer f.Close()
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, "500: Charts error.")
		return
	}
	_, err = io.Copy(f, request.Request.Body)
	defer request.Request.Body.Close()
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, "500: Charts error.")
		return
	}

	urlNew, _ := urlutil.URLJoin(pr.RepoUrl, "")
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, "500: Charts error.")
		return
	}

	log.Debug(urlNew)

	pkg.Index(pr.RepoDir, urlNew, "")

	response.WriteEntity(id)
}

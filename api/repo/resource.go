package repo

import (
	"github.com/emicklei/go-restful"
)

type RepoResourceOptions struct {
	RepoDir string
	RepoUrl string
}

//RepoResource ...
type RepoResource struct {
	RepoDir string
	RepoUrl string
}

//NewRepoResource ...
func NewRepoResource(options *RepoResourceOptions) *RepoResource {
	return &RepoResource{
		RepoDir: options.RepoDir,
		RepoUrl: options.RepoUrl,
	}
}

//Register ...
func (pr *RepoResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/envs/{env}/charts").
		Doc("Manage charts").
		Consumes(restful.MIME_XML, restful.MIME_JSON, restful.MIME_OCTET).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{chart}").To(pr.chartCtrl).
		Doc("get a chart").
		Operation("chartCtrl").
		Param(ws.PathParameter("chart", "identifier of the chart file").DataType("string")))

	ws.Route(ws.PUT("/upload/{chart}").To(pr.uploadChartCtrl).
		Doc("upload a chart").
		Operation("uploadChartCtrl").
		Param(ws.PathParameter("chart", "identifier of the chart file").DataType("string")))

	container.Add(ws)
}

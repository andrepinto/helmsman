package api

import (
	log "github.com/sirupsen/logrus"
	"github.com/emicklei/go-restful"

	"fmt"
	"net/http"
	apiRepo "github.com/andrepinto/helmsman/api/repo"
	"github.com/emicklei/go-restful-swagger12"
)

type ApiServerOptions struct {
	Port int
	RepoDir string
	RepoUrl string
}

type ApiServer struct {
	Port int
	ApiServerRepo *ApiServerRepo
}

type ApiServerRepo struct {
	RepoDir string
	RepoUrl string
}

func NewApiServer(options *ApiServerOptions) *ApiServer{
	return &ApiServer{
		Port: options.Port,
		ApiServerRepo: &ApiServerRepo{
			RepoUrl: options.RepoUrl,
			RepoDir: options.RepoDir,
		},
	}
}

func (sv *ApiServer) Run() error{

	log.Debug("Starting Http Server")

	wsContainer := restful.NewContainer()

	repoResource:= apiRepo.NewRepoResource(&apiRepo.RepoResourceOptions{
		RepoUrl: sv.ApiServerRepo.RepoUrl,
		RepoDir: sv.ApiServerRepo.RepoDir,
	})

	repoResource.Register(wsContainer)

	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(),
		WebServicesUrl: fmt.Sprintf("localhost:%d", sv.Port),
		ApiPath:        "/apidocs.json",
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "./node_modules/swagger-ui/dist"}
	swagger.RegisterSwaggerService(config, wsContainer)

	server := &http.Server{Addr: fmt.Sprintf(":%d", sv.Port), Handler: wsContainer}

	err := server.ListenAndServe()

	return err
}



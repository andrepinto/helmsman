package api

import (
	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"

	"fmt"
	"net/http"

	apiRepo "github.com/andrepinto/helmsman/api/repo"
	"github.com/emicklei/go-restful-swagger12"
)

//ServerOptions ...
type ServerOptions struct {
	Port    int
	RepoDir string
	RepoUrl string
}

//Server ...
type Server struct {
	Port       int
	ServerRepo *ServerRepo
}

//ServerRepo ...
type ServerRepo struct {
	RepoDir string
	RepoUrl string
}

//NewServer ...
func NewServer(options *ServerOptions) *Server {
	return &Server{
		Port: options.Port,
		ServerRepo: &ServerRepo{
			RepoUrl: options.RepoUrl,
			RepoDir: options.RepoDir,
		},
	}
}

//Run ...
func (sv *Server) Run() error {

	log.Debug("Starting Http Server")

	log.Debug(sv.ServerRepo)

	wsContainer := restful.NewContainer()

	repoResource := apiRepo.NewRepoResource(&apiRepo.RepoResourceOptions{
		RepoUrl: sv.ServerRepo.RepoUrl,
		RepoDir: sv.ServerRepo.RepoDir,
	})

	repoResource.Register(wsContainer)

	config := swagger.Config{
		WebServices:     wsContainer.RegisteredWebServices(),
		WebServicesUrl:  fmt.Sprintf("localhost:%d", sv.Port),
		ApiPath:         "/apidocs.json",
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "./node_modules/swagger-ui/dist"}
	swagger.RegisterSwaggerService(config, wsContainer)

	server := &http.Server{Addr: fmt.Sprintf(":%d", sv.Port), Handler: wsContainer}

	err := server.ListenAndServe()

	return err
}

package server

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	common "github.com/nivesh-star/ondc/src/common"
)

type Server struct {
	conf *common.RuntimeConfiguration
}

func NewServer() *Server {
	return &Server{
		conf: &common.RuntimeConfiguration{ApiPort: 4000},
	}
}

const (
	apiRoot = "/ondc/api/v1"
)

func setUpRouter() *gin.Engine {

	router := gin.New()

	router.Use(
		//gin.LoggerWithWriter(gin.DefaultWriter, "/ping"),
		gin.Recovery(),
	)

	api := router.Group(apiRoot)

	for _, route := range Routes {
		api.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
	return router
}

func (s *Server) Run() {

	godotenv.Load(".devenv")
	appRouter := setUpRouter()

	apiServer, _ := common.SetupHTTPServer(s.conf.ApiPort, appRouter)

	go common.RunHTTPServer("ONDC API Server (niveshStar)", apiServer)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	<-quit // wait until signal detected
	slog.Info("Server shutdown triggered")
}

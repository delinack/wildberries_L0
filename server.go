package L0

import (
	"L0/pkg/handler"
	"L0/pkg/service"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server

	natsService *service.NatsService
}

func NewServer(
	httpPort string,
	handler *handler.Handler,
	natsService *service.NatsService,
) *Server {
	return &Server{
		httpServer:  newHttpServer(httpPort, handler.InitRoutes()),
		natsService: natsService,
	}
}

func (s *Server) Run() error {
	go s.natsService.NatsServer.Subscribe(s.natsService.ClusterID, s.natsService.CreateOrder)
	s.httpServer.ListenAndServe()

	return nil
}

func newHttpServer(port string, handler http.Handler) *http.Server {
	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return httpServer
}

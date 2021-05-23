package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imatheus-lucas/golang_api/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "3333",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("server is runnig at port:", s.port)
	log.Fatal(router.Run(":" + s.port))
}

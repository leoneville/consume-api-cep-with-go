package server

import (
	"cep/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type SERVER struct {
	port   string
	server *gin.Engine
}

func Server() SERVER {
	return SERVER{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *SERVER) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}

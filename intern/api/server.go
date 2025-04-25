package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/siddheshRajendraNimbalkar/intern/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/intern/middleware"
	"github.com/siddheshRajendraNimbalkar/intern/token"
	"github.com/siddheshRajendraNimbalkar/intern/util"
)

type Server struct {
	config     util.Config
	store      *db.SQLStore
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, store *db.SQLStore) *Server {
	server := &Server{
		store: store,
	}

	tokenMaker, err := token.NewPasetoMaker(config.Secret)
	if err != nil {
		panic(err)
	}
	server.tokenMaker = tokenMaker
	router := gin.Default()

	// User
	router.POST("/user/sign-up", server.registerUser)
	router.POST("/user/sign-in", server.loginUser)

	// Patient
	authRou := router.Group("/patient")
	authRou.Use(middleware.AuthMiddleware(tokenMaker))
	authRou.POST("/create", server.CreatePatient)
	authRou.GET("/get/:id", server.GetPatient)
	authRou.GET("/get-all", server.ListPatients)
	authRou.PUT("/update/:id", server.UpdatePatient)
	authRou.DELETE("/delete/:id", server.DeletePatient)

	server.router = router
	server.config = config

	return server
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

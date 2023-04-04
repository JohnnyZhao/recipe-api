package app

import "github.com/gin-gonic/gin"

type Server struct {
	api    *Api
	router *gin.Engine
}

func NewServer(api *Api) *Server {
	router := gin.Default()
	return &Server{api: api, router: router}
}

func (s *Server) Setup() {
	s.router.POST("/recipes", s.api.CreateRecipe)
	s.router.GET("/recipes", s.api.ListRecipes)
	s.router.GET("/recipes/:id", s.api.GetByID)
	s.router.PATCH("/recipes/:id", s.api.PatchByID)
	s.router.DELETE("/recipes/:id", s.api.DeleteByID)
}

func (s *Server) Run(addr string) {
	s.router.Run(addr)
}

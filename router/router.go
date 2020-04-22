package router

import (
	"apiserver/handler/movie"
	"apiserver/handler/oss"
	"apiserver/handler/sd"
	"apiserver/handler/upload"
	"apiserver/handler/user"
	"apiserver/handler/userprofile"
	"apiserver/handler/websockets"
	"apiserver/router/middleware"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"net/http"
)

func Load(g *gin.Engine, m *melody.Melody, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	// api for authentication functionalities
	g.POST("/login", user.Login)
	g.GET("/logout", middleware.AuthMiddleware(), user.Logout)
	g.GET("/refresh", user.Refresh)
	//api get AliyunOss signature
	g.GET("/oss/generatesignature", oss.WebUploadSign)

	apiV1 := g.Group("/v1")
	{
		//user
		apiV1.POST("/user/create", user.Create)
		apiV1.DELETE("/user/delete", middleware.AuthMiddleware(), user.Delete)
		apiV1.GET("/user/get", middleware.AuthMiddleware(), user.Get)
		apiV1.GET("/user/list", user.List)
		apiV1.PUT("/user/update", middleware.AuthMiddleware(), user.Update)

		//user_profile
		apiV1.POST("/userProfile/upsert", middleware.AuthMiddleware(), userprofile.Upsert)

		//upload
		apiV1.POST("/upload", upload.Upload)
		apiV1.POST("/oss/upload", oss.Upload)

		//movie
		apiV1.POST("/movie/create", movie.Create)
		apiV1.PUT("/movie/update", movie.Update)
		apiV1.GET("/movie/get", movie.Get)
		apiV1.GET("/movie/list", movie.List)
		apiV1.DELETE("/movie/delete", movie.Delete)
	}

	//The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/demo1", sd.DemoOne)
		svcd.GET("/demo2", sd.DemoTwo)
		svcd.GET("/demo3", sd.DemoThree)
	}

	//public static
	publicRouter := g.Group("/public")
	{
		publicRouter.Static("", "public")
	}

	g.GET("/ws", func(c *gin.Context) {
		websockets.Ws(c, m)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		websockets.Message(m, s, msg)
	})

	m.HandleConnect(websockets.Connect)

	m.HandleDisconnect(websockets.Disconnect)
	return g
}

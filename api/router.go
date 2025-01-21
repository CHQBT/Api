package api

import (
	_ "milliy/api/docs"
	"milliy/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @title All
// @version 1.0
// @description API Gateway
// BasePath: /
func Router(h handler.HandlerInterface) *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	vers := router.Group("/v1")
	user := vers.Group("/auth")
	{
		user.POST("/login", h.UserMethods().Login)
	}
	twit := vers.Group("/twit")
	{
		twit.POST("", h.EnforcerMethods().CheckPermissionMiddleware(), h.TwitMethods().CreateTwit)
		twit.GET("/:id", h.TwitMethods().GetTwit)
		twit.DELETE("/:id", h.EnforcerMethods().CheckPermissionMiddleware(), h.TwitMethods().DeleteTwit)
		twit.POST("/:id", h.TwitMethods().AddReadersCount)
		twit.GET("/all", h.TwitMethods().GetAllTwits)
		twit.GET("/type/:type", h.TwitMethods().GetTwitsByType)
		twit.GET("/most-viewed", h.TwitMethods().GetMostViewedTwits)
		twit.GET("/latest-uploaded", h.TwitMethods().GetLatestTwits)
		twit.GET("/search", h.TwitMethods().SearchTwit)
		twit.POST("/location", h.EnforcerMethods().CheckPermissionMiddleware(), h.TwitMethods().CreateLocation)
		twit.POST("/url", h.EnforcerMethods().CheckPermissionMiddleware(), h.TwitMethods().CreateUrl)
		twit.POST("/photo/:twit_id", h.EnforcerMethods().CheckPermissionMiddleware(), h.TwitMethods().CreatePhoto)
		twit.POST("/video/:twit_id", h.EnforcerMethods().CheckPermissionMiddleware(), h.TwitMethods().CreateVideo)
		twit.POST("/music/:twit_id", h.EnforcerMethods().CheckPermissionMiddleware(), h.TwitMethods().CreateMusic)
	}

	return router
}

package routers

import (
	"gin-tutorial/middleware/jwt"
	"gin-tutorial/pkg/settings"
	"gin-tutorial/routers/api"
	v1 "gin-tutorial/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(settings.RunMode)

	// Test router
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "test",
		})
	})

	// Auth router
	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1/")

	// Jwt middleware
	apiv1.Use(jwt.JWT())
	{
		// Tag routers
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// Article routers
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticles)
	}
	return r
}

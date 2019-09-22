package routers

import (
	"labac/middleware/jwt"
	"labac/pkg/setting"
	"labac/routers/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	gin.SetMode(setting.RunMode)

	r.GET("/login", api.Login)

	hander := r.Group("/api")
	{
		hander.GET("/paper", api.GetPaperList)
		hander.GET("/paper/download/:name/*type", api.DownloadFile)
		hander.POST("/paper/uploadfiles", api.UploadFiles)
	}
	admin := r.Group("/admin")
	admin.Use(jwt.JWT())
	{
		admin.DELETE("/paper/:name", api.DeletePaper)
		admin.POST("/paper", api.AddPaper)
	}
	return r
}

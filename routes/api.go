package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/apis/v1"
	"github.com/jangozw/gin-api-common/middlewares"
	"github.com/jangozw/gin-api-common/utils"
	"time"
)

func RegisterRouters(router *gin.Engine) *gin.Engine {
	router.Use(middlewares.CommonMiddleware, middlewares.LoggerToFile())
	registerNoLogin(router)
	registerV1(router)
	return router
}
func registerV1(router *gin.Engine) {
	router.Group("/v1", middlewares.ApiMiddleware).
		POST("/logout", v1.Logout).
		GET("/user/list", v1.UserList).
		GET("/user/detail", v1.UserDetail)
}

func registerNoLogin(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		utils.Ctx(c).Success(map[string]string{
			"a": "Welcome! " + time.Now().Format(utils.YMDHIS),
		})
	})
	router.POST("/user/add", v1.AddUser)
	router.POST("/login", v1.Login)
}

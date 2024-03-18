package route

import (
	"forum/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	return r
}

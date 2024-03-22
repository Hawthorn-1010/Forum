package route

import (
	"forum/controller"
	"forum/middleware"
	"forum/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/ping", func(c *gin.Context) {
			userID, ok := c.Get("userID")
			if !ok {
				response.Fail(c, response.CodeServerBusy)
			}
			response.Success(c, userID)
		})
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.PostDetailHandler)
		v1.GET("/post", controller.PostListHandler)

		v1.GET("/post2", controller.PostList2Handler)

		v1.POST("/vote", controller.VoteHandler)
	}

	// 处理请求路由不存在的情况
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}

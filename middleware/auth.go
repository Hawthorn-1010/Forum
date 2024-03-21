package middleware

import (
	"forum/pkg/jwt"
	"forum/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			// 不符合格式
			response.Fail(c, response.CodeInvalidAuthFormat)
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			response.Fail(c, response.CodeInvalidToken)
			c.Abort()
			return
		}
		// 验证通过，获取claims中userId
		//username := claims.Username
		//userId := claims.UserID
		//var user = &model.User{
		//	Username: username,
		//	UserID:   userId,
		//}
		//user, err = database.QueryUser(user)

		//// 如果用户不存在
		//if user.ID == 0 {
		//	c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "权限不足！"})
		//	c.Abort()
		//	return
		//}

		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userID", claims.UserID)

		c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
	}
}

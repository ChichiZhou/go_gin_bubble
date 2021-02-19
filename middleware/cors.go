package middleware

// import (
// 	"time"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		cors.New(cors.Config{
// 			AllowOrigins:  []string{"*"}, // 允许所有域名跨域访问
// 			AllowMethods:  []string{"*"},
// 			AllowHeaders:  []string{"Origin"},
// 			ExposeHeaders: []string{"Content-Length"},
// 			// AllowCredentials: true,
// 			// AllowOriginFunc: func(origin string) bool {
// 			// 	return origin == "https://github.com"
// 			// },
// 			MaxAge: 12 * time.Hour,
// 		})
// 	}
// }

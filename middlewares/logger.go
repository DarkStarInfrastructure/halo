package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s | %s | %s | %s | %s | %d | %s | %s | %s \n",
			param.TimeStamp.Format("2006-01-02 15:04:05.000"),
			param.ClientIP,
			param.Request.Proto,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Request.UserAgent(),
			param.Latency,
			param.ErrorMessage,
		)
	})
}

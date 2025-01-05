package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogURI:       true,
		LogError:     true,
		LogMethod:    true,
		LogUserAgent: true,
		LogLatency:   true,
		HandleError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				// example log format
				// [2021-09-01 00:00:00] method=GET uri=/api/v1/users status=500 error=internal server error duration=1.2345ms UserAgent=Mozilla/5.0
				fmt.Printf("[%v] method=%v uri=%v, status=%v error='%v' duration=%v UserAgent=%v\n", time.Now().Format("2006-01-02 15:04:05"), v.Method, v.URI, v.Status, v.Error, v.Latency, v.UserAgent)
			} else {
				// example log format
				// [2021-09-01 00:00:00] method=GET uri=/api/v1/users status=200 duration=1.2345ms UserAgent=Mozilla/5.0
				fmt.Printf("[%v] method=%v uri=%v, status=%v duration=%v UserAgent=%v\n", time.Now().Format("2006-01-02 15:04:05"), v.Method, v.URI, v.Status, v.Latency, v.UserAgent)
			}
			return nil
		},
	})
}

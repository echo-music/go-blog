package middleware

import (
	"github.com/echo-music/go-blog/pkg/known"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{http.MethodGet, http.MethodDelete, http.MethodHead, http.MethodOptions, http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodConnect, http.MethodTrace},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Cookie", "Authorization", known.XRequestIDKey, "X-Auth-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

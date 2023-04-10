package middleware

import (
	"errors"
	"github.com/echo-music/go-blog/pkg/known"
	"github.com/echo-music/go-blog/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func Csrf() gin.HandlerFunc {
	return func(c *gin.Context) {
		//取cookie的csrf的随机值
		cookieToken := ""
		cookie, err := c.Request.Cookie(known.CsrfKey)
		if err == nil {
			cookieToken = cookie.Value
		}
		//未取到该值重新设置
		if cookieToken == "" {
			tokenUid, err := uuid.NewUUID()
			if err != nil {
				response.Error(c, err)
				return
			}
			cookieToken = tokenUid.String()
		}

		switch c.Request.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:

		default:
			token := c.Request.Header.Get(known.TokenRequestKey)
			if token == "" {
				token = c.GetHeader(known.TokenRequestKey)
			}

			if !strings.EqualFold(cookieToken, token) {
				c.AbortWithError(http.StatusForbidden, errors.New("invalid CSRF token"))
				return
			}
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     known.CsrfKey,
			Value:    cookieToken,
			Path:     "/",
			Expires:  time.Now().Add(12 * time.Hour),
			HttpOnly: true,
		})

		c.Next()

	}
}

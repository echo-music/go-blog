package middleware

import (
	"errors"
	"fmt"
	"github.com/echo-music/go-blog/pkg/util/grand"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// Config is the configuration struct for CSRF feature.
type Config struct {
	TokenLength     int
	TokenRequestKey string
	ExpireTime      time.Duration
	Cookie          *http.Cookie
}

var (
	// DefaultCSRFConfig is the default CSRF middleware config.
	DefaultCSRFConfig = Config{
		Cookie: &http.Cookie{
			Name: "_csrf",
		},
		ExpireTime:      0,
		TokenLength:     32,
		TokenRequestKey: "X-CSRF-Token",
	}
)

func Csrf() gin.HandlerFunc {
	return NewWithCfg(DefaultCSRFConfig)
}

func NewWithCfg(cfg Config) gin.HandlerFunc {
	return func(r *gin.Context) {
		// Read the token in the request cookie
		tokenInCookie, err := r.Request.Cookie(cfg.Cookie.Name)
		if err != nil {

		}
		tokenInCookieStr := tokenInCookie.Value
		fmt.Println(tokenInCookieStr)
		if tokenInCookie.String() == "" {
			// Generate a random token
			token, err := grand.S(cfg.TokenLength)
			if err != nil {
				r.Error(err)
				r.Abort()
				return
			}
			tokenInCookieStr = string(token)
		}

		// Read the token attached to the request
		// Read priority: Router < Query < Body < Form < Custom < Header
		tokenInRequestData := r.GetHeader(cfg.TokenRequestKey)
		if tokenInRequestData == "" {
			tokenInRequestData = r.Request.Header.Get(cfg.TokenRequestKey)
		}

		switch r.Request.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
			// No verification required
		default:
			// Authentication token
			if !strings.EqualFold(tokenInCookieStr, tokenInRequestData) {
				r.AbortWithError(http.StatusForbidden, errors.New("invalid CSRF token"))
				return
			}
		}

		// Set cookie timeout
		cfg.Cookie.Expires = time.Now().Add(cfg.ExpireTime)
		cfg.Cookie.Value = tokenInCookieStr

		// Set cookie in response
		http.SetCookie(r.Writer, cfg.Cookie)
		r.Next()
	}
}

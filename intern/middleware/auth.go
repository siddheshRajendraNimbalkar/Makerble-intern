package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/siddheshRajendraNimbalkar/intern/token"
)

const (
	AuthHeaderKey  = "Authorization"
	AuthTypeBearer = "bearer"
	AuthPayloadKey = "auth_payload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(AuthHeaderKey)
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			ctx.Abort()
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) != 2 || strings.ToLower(fields[0]) != AuthTypeBearer {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be 'Bearer <token>'"})
			ctx.Abort()
			return
		}

		tokenStr := fields[1]
		payload, err := tokenMaker.VerifyToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Set(AuthPayloadKey, payload)
		ctx.Next()
	}
}

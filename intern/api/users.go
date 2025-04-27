package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/siddheshRajendraNimbalkar/intern/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/intern/token"
	"golang.org/x/crypto/bcrypt"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=receptionist doctor"`
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	AccessToken string    `json:"access_token"`
	UserId      int64     `json:"user_id"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func (server *Server) registerUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while hashing password",
			"error":   err.Error(),
		})
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: string(hashPassword),
		Role:     req.Role,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	maker, err := token.NewPasetoMaker(server.config.Secret)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tokenStr, payload, err := maker.CreateToken(user.Username, int64(user.ID), user.Role, server.config.JwtDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating token",
			"error":   err.Error(),
		})
		return
	}

	resp := loginResponse{
		AccessToken: tokenStr,
		UserId:      int64(user.ID),
		ExpiresAt:   payload.ExpiresAt,
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := server.store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong password",
			"error":   err.Error(),
		})
		return
	}

	tokenString, payload, err := server.tokenMaker.CreateToken(user.Username, int64(user.ID), user.Role, server.config.JwtDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	resp := loginResponse{
		AccessToken: tokenString,
		UserId:      int64(user.ID),
		ExpiresAt:   payload.ExpiresAt,
	}

	ctx.JSON(http.StatusOK, resp)
}

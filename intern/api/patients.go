package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/siddheshRajendraNimbalkar/intern/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/intern/token"
)

type createPatientRequest struct {
	FullName string `json:"full_name" binding:"required,min=2,max=100"`
	Age      int32  `json:"age" binding:"required,gte=0,lte=150"`
	Gender   string `json:"gender" binding:"required,oneof=male female other"`
	Address  string `json:"address" binding:"omitempty,max=200"`
	Phone    string `json:"phone" binding:"omitempty,e164"`
	Status   string `json:"status" binding:"required,oneof=admitted discharged 'under observation'"`
}

func (server *Server) CreatePatient(ctx *gin.Context) {

	var req createPatientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}
	authPaylod := ctx.MustGet("auth_payload").(*token.Payload)

	arg := db.CreatePatientParams{
		FullName:  req.FullName,
		Age:       req.Age,
		Gender:    req.Gender,
		Address:   sql.NullString{String: req.Address, Valid: req.Address != ""},
		Phone:     sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Status:    req.Status,
		CreatedBy: sql.NullInt32{Int32: int32(authPaylod.UserID), Valid: true},
	}

	patient, err := server.store.CreatePatient(ctx, arg)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to create patient", "details": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"patient": patient})
}

func (server *Server) GetPatient(ctx *gin.Context) {
	var req struct {
		ID int32 `uri:"id" binding:"required,min=1"`
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient, err := server.store.GetPatient(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	ctx.JSON(http.StatusOK, patient)
}

func (server *Server) ListPatients(ctx *gin.Context) {
	patients, err := server.store.ListPatients(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, patients)
}

type updatePatientRequest struct {
	FullName string `json:"full_name" binding:"required,min=2"`
	Age      int32  `json:"age" binding:"required,gte=0"`
	Gender   string `json:"gender" binding:"required,oneof=male female other"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Status   string `json:"status" binding:"required,oneof=admitted discharged 'under observation'"`
}

func (server *Server) UpdatePatient(ctx *gin.Context) {
	idStr := ctx.Params.ByName("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a valid integer"})
		return
	}

	var req updatePatientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdatePatientParams{
		ID:       int32(id),
		FullName: req.FullName,
		Age:      req.Age,
		Gender:   req.Gender,
		Address:  sql.NullString{String: req.Address, Valid: req.Address != ""},
		Phone:    sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Status:   req.Status,
	}

	patient, err := server.store.UpdatePatient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, patient)
}

func (server *Server) DeletePatient(ctx *gin.Context) {
	var req struct {
		ID int32 `uri:"id" binding:"required,min=1"`
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := server.store.DeletePatient(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

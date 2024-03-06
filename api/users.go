package api

import (
	"database/sql"
	db "moneytracker/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	req := createUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
	}

	user, err := server.queries.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetUser(ctx *gin.Context) {
	req := getUserRequest{}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.queries.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListUsers(ctx *gin.Context) {
	req := listUsersRequest{}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.queries.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type updateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) UpdateUser(ctx *gin.Context) {
	loc := getUserRequest{}
	req := updateUserRequest{}

	if err := ctx.ShouldBindUri(&loc); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		Name: req.Name,
		ID:   loc.ID,
	}

	user, err := server.queries.UpdateUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) DeleteUser(ctx *gin.Context) {
	loc := getUserRequest{}

	if err := ctx.ShouldBindUri(&loc); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.queries.DeleteUser(ctx, loc.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully id " + strconv.Itoa(int(loc.ID))})
}

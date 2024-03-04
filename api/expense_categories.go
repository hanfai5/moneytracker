package api

import (
	"database/sql"
	db "moneytracker/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createExpenseCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

func (server *Server) CreateExpenseCategory(ctx *gin.Context) {

	req := createExpenseCategoryRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateExpenseCategoryParams{
		Name:  req.Name,
		Color: req.Color,
	}

	category, err := server.queries.CreateExpenseCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type getExpenseCategoryRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetExpenseCategory(ctx *gin.Context) {

	req := getExpenseCategoryRequest{}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	category, err := server.queries.GetExpenseCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type listExpenseCategoriesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListExpenseCategories(ctx *gin.Context) {
	req := listExpenseCategoriesRequest{}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListExpenseCategoriesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	categories, err := server.queries.ListExpenseCategories(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

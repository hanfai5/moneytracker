package api

import (
	"database/sql"
	"fmt"
	db "moneytracker/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createIncomeCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

func (server *Server) CreateIncomeCategory(ctx *gin.Context) {

	req := createIncomeCategoryRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateIncomeCategoryParams{
		Name:  req.Name,
		Color: req.Color,
	}

	category, err := server.queries.CreateIncomeCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type getIncomeCategoryRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetIncomeCategory(ctx *gin.Context) {

	req := getIncomeCategoryRequest{}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.queries.GetIncomeCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type listIncomeCategoriesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListIncomeCategories(ctx *gin.Context) {
	req := listIncomeCategoriesRequest{}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)

	arg := db.ListIncomeCategoriesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	categories, err := server.queries.ListIncomeCategories(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

type updateIncomeCategoryNameRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) UpdateIncomeCategoryName(ctx *gin.Context) {

	loc := getIncomeCategoryRequest{}
	req := updateIncomeCategoryNameRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&loc); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateIncomeCategoryNameParams{
		Name: req.Name,
		ID:   loc.ID,
	}

	category, err := server.queries.UpdateIncomeCategoryName(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type updateIncomeCategoryColorRequest struct {
	Color string `json:"color" binding:"required"`
}

func (server *Server) UpdateIncomeCategoryColor(ctx *gin.Context) {
	loc := getIncomeCategoryRequest{}
	req := updateIncomeCategoryColorRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&loc); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateIncomeCategoryColorParams{
		Color: req.Color,
		ID:    loc.ID,
	}

	category, err := server.queries.UpdateIncomeCategoryColor(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (server *Server) DeleteIncomeCategory(ctx *gin.Context) {
	loc := getIncomeCategoryRequest{}

	if err := ctx.ShouldBindUri(&loc); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.queries.DeleteIncomeCategory(ctx, loc.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully id " + strconv.Itoa(int(loc.ID))})
}

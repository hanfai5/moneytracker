package api

import (
	db "moneytracker/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP service for our money tracking service
type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(queries *db.Queries) *Server {
	server := Server{queries: queries}
	router := gin.Default()

	router.POST("/income_categories", server.CreateIncomeCategory)
	router.GET("/income_categories/:id", server.GetIncomeCategory)
	router.GET("/income_categories", server.ListIncomeCategories)
	router.PATCH("/income_categories/name/:id", server.UpdateIncomeCategoryName)
	router.PATCH("/income_categories/color/:id", server.UpdateIncomeCategoryColor)
	router.DELETE("/income_categories/:id", server.DeleteIncomeCategory)
	router.POST("/expense_categories", server.CreateExpenseCategory)
	router.GET("/expense_categories/:id", server.GetExpenseCategory)
	router.GET("/expense_categories", server.ListExpenseCategories)
	server.router = router
	return &server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

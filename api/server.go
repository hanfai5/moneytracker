package api

import (
	"fmt"
	db "moneytracker/db/sqlc"
	"moneytracker/db/util"
	"moneytracker/token"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP service for our money tracking service
type Server struct {
	config     util.Config
	queries    *db.Queries
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, queries *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := Server{
		config:     config,
		queries:    queries,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return &server, nil
}

func (server *Server) setupRouter() {
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
	router.PATCH("/expense_categories/color/:id", server.UpdateExpenseCategoryColor)
	router.PATCH("/expense_categories/name/:id", server.UpdateExpenseCategoryName)
	router.DELETE("/expense_categories/:id", server.DeleteExpenseCategory)
	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.LoginUser)
	router.GET("/users/:id", server.GetUser)
	router.GET("/users", server.ListUsers)
	router.PATCH("/users/:id", server.UpdateUser)
	router.DELETE("/users/:id", server.DeleteUser)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

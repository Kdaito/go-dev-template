package server

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Kdaito/go-dev-template/internal/application/handler"
	"github.com/Kdaito/go-dev-template/internal/application/middleware"
	"github.com/Kdaito/go-dev-template/internal/domain/service"
	"github.com/Kdaito/go-dev-template/internal/infrastructure"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Server struct {
	db     *sql.DB
	router *echo.Echo
}

func NewServer() *Server {
	return &Server{}
}

// setUpDb sets up a connection to the database.
func (server *Server) setUpDb() {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_ROOT_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")

	var datasource = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)

	if db, err := sql.Open("mysql", datasource); err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	} else {
		server.db = db
		fmt.Println("db connected!!")
	}
}

// setUpRouter sets up the server's router.
func (server *Server) setUpRouter() {
	router := echo.New()

	// middleware
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CorsMiddleware())

	// DI
	userRepo := infrastructure.NewUser(server.db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	v1 := router.Group("/v1")

	// routing
	v1.GET("/users", userHandler.GetUserList)
	v1.GET("/users/:id", userHandler.GetUserByID)
	v1.POST("/users", userHandler.CreateUser)

	server.router = router
}

// Start starts the server.
func (server *Server) Start() {
	server.setUpDb()
	server.setUpRouter()

	if err := server.router.Start(":8080"); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}

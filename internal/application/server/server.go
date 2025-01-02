package server

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	db *sql.DB
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{}
}

// setUpDb sets up a connection to the database.
func (server *Server) setUpDb() {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_ROOT_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")

	fmt.Printf("user: %s, pw: %s, db_name: %s", user, pw, db_name)

	var path = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)

	if db, err := sql.Open("mysql", path)	; err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	} else {
		server.db = db
		fmt.Println("db connected!!")
	}
}

// setUpRouter sets up the server's router.
func (server *Server) setUpRouter() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongpongpong",
		})
	})

	server.router = router
}

// Start starts the server.
func (server *Server) Start() {
	server.setUpDb()
	server.setUpRouter()

	if err := server.router.Run(":8080"); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}

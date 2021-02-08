package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dongil91/module-test/mysql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	connString := "go_api:go_api@tcp(172.18.0.2:3306)/go_api"
	db, err := setupDatabase(connString)
	if err != nil {
		return err
	}

	userRepository := mysql.NewMysqlUserRepository(db)
	router := gin.Default()
	router.GET("/apis/users/me", func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		log.Print(authorization)
		me, _ := userRepository.FindById(c, "test")
		response := make(map[string]string)
		response["message"] = "success"
		c.JSON(http.StatusOK, me)
	})
	router.Run()

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

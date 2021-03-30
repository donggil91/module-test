package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dongil91/module-test/repository"
	"github.com/dongil91/module-test/router"
	"github.com/dongil91/module-test/service"
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
	connString := "go_api:go_api@tcp(172.18.0.2:3306)/go_api?parseTime=true"
	db, err := setupDatabase(connString)
	if err != nil {
		return err
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	userRepository := repository.NewMysqlSellerRepository(db)
	userService := service.NewUserService(userRepository)
	engine := gin.Default()
	router.NewSellerRouter(engine, userService)
	engine.Run()

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

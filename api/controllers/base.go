package controllers

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	dsn := "host=localhost user=root password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	//fmt.Printf("We are connected to the %s database", Dbdriver)
	// dsn := "host=localhost  user=root password=admin dbname=fp_local port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	// server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{
	// 	SkipDefaultTransaction: true,
	// 	PrepareStmt:            true,
	// })

	// server.DB, err := gorm.Open(sqlserver.Open(DBURL), &gorm.Config{})
	//gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", Dbdriver)
	}

	//server.DB.AutoMigrate(&models.RegisteredDevices{})

	// server.DB.Debug().AutoMigrate(models.Login{}) //database migration

	server.Router = gin.New()
	server.Router.Use(gin.Recovery(), gin.Logger())

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 9072")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

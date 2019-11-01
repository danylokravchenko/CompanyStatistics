package main

import (
	"fmt"
	"github.com/UndeadBigUnicorn/CompanyStatistics/dbworker"
	"github.com/UndeadBigUnicorn/CompanyStatistics/handlers"
	"github.com/UndeadBigUnicorn/CompanyStatistics/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
)

func main() {

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//gin.SetMode(gin.ReleaseMode)

	route := gin.Default()

	route.Use(middleware.Auth())

	go func() {

		company := route.Group("/company")
		{
			company.POST("/add", handlers.AddCompany)
			company.GET("/stats", handlers.GetStats)
			company.POST("/update", handlers.UpdateCompany)
		}

		route.Run(":8080")

	}()


	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	dbworker.CloseConnection()
	fmt.Println("Closing connection to database")
	fmt.Println("End of a program")

}

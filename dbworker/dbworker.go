package dbworker

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var db *sqlx.DB


// Initialize database connection
func init() {

	dataSource := os.Getenv("cr_mysql_uri")

	var err error
	db, err = sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Fatalln(err)
	}

}


// Close connection to database
func CloseConnection() error {
	return db.Close()
}

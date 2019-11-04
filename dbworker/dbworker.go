package dbworker

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB


// Initialize database connection
func init() {

	var err error
	db, err = sqlx.Connect("mysql", config.GetSetting("mysqlURL").(string))
	if err != nil {
		log.Fatalln(err)
	}

}


// Close connection to database
func CloseConnection() error {
	return db.Close()
}

package dbworker

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/config"
	. "github.com/UndeadBigUnicorn/CompanyStatistics/infrastructure/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB


// Initialize database connection
func init() {

	var err error
	db, err = sqlx.Connect("mysql", config.GetSetting("mysqlURL").(string))
	if err != nil {
		Error.Fatalln(err)
	}

}


// Close connection to database
func CloseConnection() error {
	return db.Close()
}

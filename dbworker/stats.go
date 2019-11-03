package dbworker

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/infrastructure/timeParser"
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
)

// Load companies from database
func LoadStats(companyID uint64) []models.UserStats {

	var users []models.UserStats
	db.Select(&users, `
		select
			u.id, concat(u.firstname, ' ', u.lastname) as name, s.today, 
			ifnull(sum(s.created),0) as created, ifnull(sum(s.opened),0) as opened
		from stats s
		inner join users u on u.id = s.userid
		where u.deletedat is null and u.companyid = ?
	`, &companyID)


	for idx, user := range users {
		user.TimeToday = timeParser.ParseTime(user.Today)
		users[idx] = user
	}

	return users

}

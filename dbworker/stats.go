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
			s.id as statsid, u.id, concat(u.firstname, ' ', u.lastname) as name, s.today, 
			ifnull(sum(s.created),0) as created, ifnull(sum(s.opened),0) as opened
		from stats s
		inner join users u on u.id = s.userid
		where u.deletedat is null and u.companyid = ?
		group by s.today
	`, &companyID)

	for idx, user := range users {
		user.TimeToday = timeParser.ParseTime(user.Today)
		users[idx] = user
	}

	return users

}


// Wrap stats insert queries in a transaction
func InsertBatchStats(stats []models.UserStats) ([]models.UserStats, error) {

	tx := db.MustBegin()

	for idx, personalStats := range stats {

		res, _ := tx.NamedExec(`
		 insert into stats
                	(companyid, userid, today, created, opened)
                values
                    (:companyid, :id, :today, :created, :opened)
		`, personalStats)

		id, _ := res.LastInsertId()
		stats[idx].StatsID = uint64(id)

	}

	return stats, tx.Commit()

}


// Wrap stats update queries in a transaction
func UpdateBatchStats(stats []models.UserStats) error {

	tx := db.MustBegin()

	for _, personalStats := range stats {

		tx.NamedExec(`
			UPDATE stats
			SET 
			opened = :opened,
			created = :created
			WHERE id = :statsid;
		`, personalStats)

	}

	return tx.Commit()

}
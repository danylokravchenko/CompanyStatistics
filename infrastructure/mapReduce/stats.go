package mapReduce

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/infrastructure/sorter"
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
	"time"
)


const (
	mapSize = 30
)

// send user stats into input channel
func LoadStatsGenerateInput(users []models.UserStats) chan interface{} {

	input := make(chan interface{})

	go func() {
		begin, end := 0, 0
		l := len(users)

		// split 1 big array of UserStats into small arrays
		// and send into mapper channel
		for end + mapSize < l {
			end += mapSize
			input <- users[begin:end]
			begin = end
		}

		input <- users[end:]

		close(input)

	}()

	return input

}

// Mapper function for loading companies
// Transform UserStats arrays into map
func LoadStatsMapper() MapperFunc {
	return func(array interface{}, output chan interface{}) {

		// create temp chan with map
		results := make(chan models.UserStatsMap)
		statsArray := array.([]models.UserStats)
		counter := 1

		go func() {
			for _, stats := range statsArray {
				// transform array into map
				statsMap := models.UserStatsMap{}
				statsMap[stats.ID] = stats
				results <- statsMap

				if counter == len(statsArray) {
					close(results)
				}
				counter++
			}
		}()

		timeMap := models.TimeMap{}

		// add userStats map to timeMap based on time
		for res := range results {

			for id, stats := range res {

				if userStatsMap, ok := timeMap[stats.TimeToday]; !ok {
					// there are no stats on this day yet
					timeMap[stats.TimeToday] = res
				} else {
					// there are stats so we need to add data to existing map
					userStatsMap[id] = stats
				}
			}

		}

		// send temp maps into reducer input
		output <- timeMap

	}
}

// Reducer function for loading companies
// Intersect temporary timeMaps into final timeMap
func LoadStatsReducer() ReducerFunc {
	return func(input, output chan interface{}) {

		finalMap := make(models.TimeMap)

		for in := range input {
			timeMap := in.(models.TimeMap)

			for today, tempStatsMap := range timeMap {
				if userStatsMap, ok := finalMap[today]; !ok {
					// there are no stats on this day yet
					finalMap[today] = tempStatsMap
				} else {
					// add user stats on current day
					for id, stats := range tempStatsMap {
						userStatsMap[id] = stats
					}
				}
			}

		}

		output <- finalMap

	}
}


// Get detail stats for given period of time
func FilterStatsGenerateInput(timeMap models.TimeMap, dateFrom, dateTo time.Time ) chan interface{} {

	input := make(chan interface{})

	go func() {
		for today := dateFrom; today.Unix() <= dateTo.Unix(); today = today.AddDate(0,0,1) {
			if userStats, ok := timeMap[today]; ok {
				input <- userStats
			}
		}
		close(input)
	}()

	return input

}


// Intersect timeMaps, compute total stats for users
func FilterStatsMapper() MapperFunc {
	return func(tempMap interface{}, output chan interface{}) {
		//timeMap := tempMap.(models.TimeMap)
		// do nothing :)
		output <- tempMap
	}
}

// Reducer: convert maps into arrays and apply sorting by order
func FilterStatsReducer(order string) ReducerFunc {
	return func(input, output chan interface{}) {

		finalMap := make(models.UserStatsMap)

		// update total stats
		for in := range input {
			userStatsMap := in.(models.UserStatsMap)
			for id, userStats := range userStatsMap {
				if totalStats, ok := finalMap[id]; !ok {
					finalMap[id] = userStats
				} else {
					totalStats.Opened += userStats.Opened
					totalStats.Created += userStats.Created
					finalMap[id] = totalStats
				}
			}
		}

		// convert to array
		userStatsArray := make([]models.UserStats, 0)
		for _, userStats := range finalMap {
			userStatsArray = append(userStatsArray, userStats)
		}

		// sort results
		output <- sorter.SortDetailStats(userStatsArray, order)

	}
}


// Send to mapper input chanel only map with one specific company
func UpdateStatsGenerateInput(statsMap models.StatsMap) chan interface{} {

	input := make(chan interface{})

	go func() {
		for _, stats := range statsMap {
			input <- stats
		}
		close(input)
	}()

	return input

}

// Send to reducer stats (as array) that only needed an update or creation
func UpdateStatsMapper() MapperFunc {
	return func(statsInstance interface{}, output chan interface{}) {

		stats := statsInstance.(*models.Stats)
		statsArray := make([]models.UserStats, 0)

		for today, userStatsMap := range stats.TimeMap {
			for id, personalStats := range userStatsMap {
				if personalStats.UpdateIsNeeded {
					// update is not needed more
					personalStats.UpdateIsNeeded = false
					stats.TimeMap[today][id] = personalStats
					statsArray = append(statsArray, personalStats)
				}
			}
		}

		output <- statsArray

	}
}

// Append personal stats that need to be updated or created to one big array with all companies
func UpdateStatsReducer() ReducerFunc {
	return func(input chan interface{}, output chan interface{}) {

		// [0] - stats to insert to db
		// [1] - stats to update in db
		finalArray := make([][]models.UserStats, 2)

		for in := range input {
			statsArray := in.([]models.UserStats)

			for _, stats := range statsArray {
				if stats.StatsID == 0 {
					finalArray[0] = append(finalArray[0], stats)
				} else {
					finalArray[1] = append(finalArray[1], stats)
				}
			}

		}

		output <- finalArray

	}
}
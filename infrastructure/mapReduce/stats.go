package mapReduce

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
)


const (
	mapSize = 60
)

// send users into input channel
func LoadCompaniesGenerateInput(users []models.UserStats) chan interface{} {

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
func LoadCompaniesMapper() MapperFunc {
	return func(array interface{}, output chan interface{}) {

		// create temp chan with map
		results := make(chan models.UserStatsMap)
		statsArray := array.([]models.UserStats)
		counter := 0

		for _, stats := range statsArray {
			// transform array into map
			go func(stats models.UserStats) {
				statsMap := models.UserStatsMap{}
				statsMap[stats.ID] = stats
				results <- statsMap

				if counter == len(statsArray) {
					close(results)
				}
			}(stats)
			counter++
		}

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
func LoadCompaniesReducer() ReducerFunc {
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
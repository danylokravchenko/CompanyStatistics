package handlers

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/infrastructure/timeParser"
	"github.com/UndeadBigUnicorn/CompanyStatistics/services/statsService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)


// Add already created stats to cache
func AddStats(c *gin.Context) {

	var addStatsRequestModel struct {
		CompanyID string `form:"companyid" json:"companyid" binding:"required"`
		UserID string `form:"userid" json:"userid" binding:"required"`
		Name string `form:"name" json:"name" binding:"required"`
		Target string `form:"target" json:"target" binding:"required"`
		Today string `form:"today" json:"today" binding:"required"`
	}

	if err := c.ShouldBind(&addStatsRequestModel); err != nil {
		_400(c)
	}

	companyID, err := strconv.ParseUint(addStatsRequestModel.CompanyID, 10, 64)
	if err != nil || companyID <= 0 {
		_400(c)
	}

	userID, err := strconv.ParseUint(addStatsRequestModel.UserID, 10, 64)
	if err != nil || userID <= 0 {
		_400(c)
	}

	if addStatsRequestModel.Target != "created" && addStatsRequestModel.Target != "opened" {
		_400(c)
	}

	if addStatsRequestModel.Today == "" {
		_400(c)
	}

	today := timeParser.ParseTime(addStatsRequestModel.Today)

	err = statsService.UpdateStats(appCache, companyID, userID, addStatsRequestModel.Name, addStatsRequestModel.Target, today)

	c.JSON(http.StatusOK, gin.H {
		"error": err.Error(),
	})

}


// Get detail statistic for specific company for given period sorted by
func GetDetailStats (c *gin.Context) {

	var getStatsRequestModel struct {
		CompanyID string `form:"companyid" json:"companyid" binding:"required"`
		Order string `form:"order" json:"order" binding:"required"`
		From string `form:"from" json:"from" binding:"required"`
		To string `form:"to" json:"to" binding:"required"`
	}

	if err := c.ShouldBind(&getStatsRequestModel); err != nil {
		_400(c)
	}

	companyID, err := strconv.ParseUint(getStatsRequestModel.CompanyID, 10, 64)
	if err != nil || companyID <= 0 {
		_400(c)
	}

	if !strings.Contains("opened created name",getStatsRequestModel.Order){
		_400(c)
	}

	from, to := timeParser.ParseTime(getStatsRequestModel.From), timeParser.ParseTime(getStatsRequestModel.To)

	/*stats := */statsService.GetDetailStats(appCache, companyID, from, to, getStatsRequestModel.Order)

	c.JSON(http.StatusOK, gin.H {
		"error": "",
		//stats...
	})

}
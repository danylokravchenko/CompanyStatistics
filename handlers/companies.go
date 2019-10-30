package handlers

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/services/companyService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


// Get company stats
func GetStats(c *gin.Context) {

	id := c.Query("id")

	if id == "" {
		_400(c)
		return
	}

	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || companyID <= 0 {
		_400(c)
		return
	}

	company, err := appCache.GetCompany(companyID)
	if err != nil {
		_404(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"error": "",
		"id": company.ID,
		"totallocations": company.TotalLocations,
		"totaldoctors": company.TotalDoctors,
		"totalusers": company.TotalUsers,
		"totalinvitations": company.TotalInvitations,
		"totalcreatedreviews": company.TotalCreatedReviews,
		"totalopenedreviews": company.TotalOpenedReviews,
	})

}


// Add new company to cache
func AddCompany(c *gin.Context) {

	var addCompanyRequestModel struct {
		ID string `form:"id" json:"id" binding:"required"`
	}

	if err := c.ShouldBind(&addCompanyRequestModel); err != nil {
		_400(c)
		return
	}

	companyID, err := strconv.ParseUint(addCompanyRequestModel.ID, 10, 64)
	if err != nil || companyID <= 0 {
		_400(c)
		return
	}

	err = companyService.AddCompany(appCache, companyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"error": "",
	})

}
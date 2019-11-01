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


// Update company stats
func UpdateCompany(c *gin.Context) {

	var updateCompanyRequestModel struct {
		ID string `form:"id" json:"id" binding:"required"`
		TotalLocations string `form:"totallocations" json:"totallocations" binding:"required"`
		TotalDoctors string `form:"totaldoctors" json:"totaldoctors" binding:"required"`
		TotalUsers string `form:"totalusers" json:"totalusers" binding:"required"`
		TotalInvitations string `form:"totalinvitations" json:"totalinvitations" binding:"required"`
		TotalCreatedReviews string `form:"totalcreatedreviews" json:"totalcreatedreviews" binding:"required"`
		TotalOpenedReviews string `form:"totalopenedreviews" json:"totalopenedreviews" binding:"required"`
	}

	if err := c.ShouldBind(&updateCompanyRequestModel); err != nil {
		_400(c)
		return
	}

	companyID, err := strconv.ParseUint(updateCompanyRequestModel.ID, 10, 64)
	if err != nil || companyID <= 0 {
		_400(c)
		return
	}

	status, err := companyService.UpdateCompany(appCache, updateCompanyRequestModel)

	c.JSON(status, gin.H {
		"error": err.Error(),
	})

}
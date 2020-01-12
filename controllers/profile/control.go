package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/folio/core"
	"github.com/louisevanderlith/husk"
	"net/http"
	"strconv"
)

func Get(c *gin.Context) {
	results := core.GetAllProfiles(1, 10)
	c.JSON(http.StatusOK, results)
}

func Search(c *gin.Context) {
	page, size := getPageData(c.Param("pagesize"))
	hsh := c.Param("hash")

	results := core.GetProfiles(page, size, hsh)

	c.JSON(http.StatusOK, results)
}

func View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	result, err := core.GetProfile(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, result)
}

func ViewByName(c *gin.Context) {
	siteParam := c.Param("name")

	result, err := core.GetProfileByName(siteParam)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, result)
}

func Create(c *gin.Context) {
	site := core.Profile{}
	err := c.Bind(&site)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	rec := site.Create()

	if rec.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, rec.Error)
	}

	c.JSON(http.StatusOK, rec)
}

func getPageData(pageData string) (int, int) {
	defaultPage := 1
	defaultSize := 10

	if len(pageData) < 2 {
		return defaultPage, defaultSize
	}

	pChar := []rune(pageData[:1])

	if len(pChar) != 1 {
		return defaultPage, defaultSize
	}

	page := int(pChar[0]) % 32
	pageSize, err := strconv.Atoi(pageData[1:])

	if err != nil {
		return defaultPage, defaultSize
	}

	return page, pageSize
}

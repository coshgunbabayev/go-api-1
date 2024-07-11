package controller

import "github.com/gin-gonic/gin"

func GetHomePage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GetAccountPage(c *gin.Context) {
	c.HTML(200, "account.html", nil)
}

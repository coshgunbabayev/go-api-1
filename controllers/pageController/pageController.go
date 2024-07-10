package pageController

import "github.com/gin-gonic/gin"

func GetHomePage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GetSignUpPage(c *gin.Context) {
	c.HTML(200, "signup.html", nil)
}

func GetLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

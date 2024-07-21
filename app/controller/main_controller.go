package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Main website",
	})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.tmpl", gin.H{
		"title": "About page",
	})
}

func Tours(c *gin.Context) {
	c.HTML(http.StatusOK, "find.tmpl", gin.H{
		"title": "find tour page",
	})
}

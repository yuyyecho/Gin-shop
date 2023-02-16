package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowReg(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func HandleReg(c *gin.Context) {

}

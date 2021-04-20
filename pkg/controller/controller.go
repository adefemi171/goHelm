package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

)

// User struct for User details
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

var userSlice = []User{
	User{
		Username:  "Welcome",
		FirstName: "Web-Clan",
		LastName:  "Okay",
		Email:     "adefemi171@gmail.com",
	},
	User{
		Username:  "Welcome",
		FirstName: "Mugun-Clan",
		LastName:  "Okay",
		Email:     "adefemi171@gmail.com",
	},
}

// Welcome endpoint
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

// NotFound endpoint
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}

// UserHandler endpoint
func UserHandler(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	enc := json.NewEncoder(w)
	enc.SetIndent(" ", "\t")
	enc.Encode(userSlice)
}

// GetHeader to display header
func GetHeader(c *gin.Context) {
	res := gin.H{}
	for k, vals := range c.Request.Header {
		log.Infof("%s", k)
		res[k] = vals
	}
	c.JSON(200, res)
}
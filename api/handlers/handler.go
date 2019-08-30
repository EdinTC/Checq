package handlers

import (
	"net/http"

	"github.com/EdinTC/Checq/api/utils"

	"github.com/gin-gonic/gin"
)

const jsonMediaType = "application/json"
const textMediaType = "text/plain"

// Response is used for the response
type Response struct {
	IP       []string `json:"ip"`
	TXT      []string `json:"txt,omitempty"`
	NS       []string `json:"ns,omitempty"`
	Hostname string   `json:"hostname,omitempty"`
	// TODO: IP4,
}

//
// QueryHandler handles the request and returns all needed stuff
//
func QueryHandler(c *gin.Context) {
	name := c.Param("name")
	ip := utils.QueryIPAdress(name)
	txt := utils.QueryTXT(name)
	ns := utils.QueryNS(name)
	hn := utils.QueryHostname(name)

	if ip != nil {
		c.JSON(http.StatusOK, Response{
			IP:       ip,
			TXT:      txt,
			NS:       ns,
			Hostname: hn,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": Response{}})
	}
}

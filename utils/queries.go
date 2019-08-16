package utils

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//
// QueryIPAdress queries the given domain name for IP records
//
func QueryIPAdress(c *gin.Context) {
	name := c.Param("name")
	res, err := net.LookupIP(name)

	var results []string
	for _, ip := range res {
		if !stringInSlice(ip.String(), results) {
			results = append(results, ip.String())
		}
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"body": results})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//
// QueryTXT queries the given domain name for txt records.
// When www is given as a subdomain it is removed
//
func QueryTXT(c *gin.Context) {
	name := strings.TrimLeft(c.Param("name"), "www.")
	res, err := net.LookupTXT(name)

	var results []string
	for _, value := range res {
		if !stringInSlice(value, results) {
			results = append(results, value)
		}
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"body": results})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//
// QueryNS queries the given domain name for nameserver records.
//
func QueryNS(c *gin.Context) {
	name := strings.TrimLeft(c.Param("name"), "www.")
	res, err := net.LookupNS(name)

	var results []string
	for _, value := range res {
		results = append(results, value.Host)
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"body": results})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// nameserver, _ := net.LookupNS("facebook.com")
// for _, ns := range nameserver {
// 	fmt.Println(ns)
// }

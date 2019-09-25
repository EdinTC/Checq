package utils

import (
	"fmt"
	"net/url"
)

//
// StringInSlice checks if the string is already in the list
//
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

//
// ExtractDomain returns the hostname from given url
//
func ExtractDomain(requestURL string) string {
	u, err := url.Parse(requestURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)

	return ""
}

package utils

//
// This functions checks if the string is already in the list
//
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

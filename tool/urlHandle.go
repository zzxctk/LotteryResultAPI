package tool

import "strings"

func GetUrlFileName(url string) string {
	if url == "" {
		return ""
	}
	arr := strings.Split(url, "/")
	if len(arr) > 0 {
		return arr[len(arr)-1]
	} else {
		return ""
	}
}

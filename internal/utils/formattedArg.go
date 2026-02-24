package utils

import (
	"strings"
)

func FormattedArgs(strRaw string) string {
	str := strRaw
	result := strings.ReplaceAll(str, " ", "")
	return result
}

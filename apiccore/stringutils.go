package apiccore

import (
	"regexp"
	"strings"
)

//SplitByUpper split String by Upper Cases and add an space in the middle
func SplitByUpper(s string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	var result strings.Builder

	submatchall := re.FindAllString(s, -1)
	for _, element := range submatchall {
		if result.Len() > 1 {
			result.WriteString(" ")
		}
		result.WriteString(element)
	}
	return result.String()
}

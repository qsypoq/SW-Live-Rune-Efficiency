package main

import (
	"regexp"
	"strings"
)

func trimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func clean_char(to_clean string, char string) string {
	regex := regexp.MustCompile(char)
	cleaned := regex.ReplaceAllString(to_clean, "")
	return cleaned
}

func split_stats(rune_subs string) []string {
	repercent := regexp.MustCompile("\n")
	percent_done := repercent.ReplaceAllString(rune_subs, "|")
	cleaned := clean_char(percent_done, "\\)")
	cleaned = clean_char(cleaned, "\\(")
	cleaned = clean_char(cleaned, "\\(")
	cleaned = clean_char(cleaned, "©")
	result := strings.Split(cleaned, "|")
	// for i := range result {
	// 	fmt.Print("\n", result[i])
	// }
	return result
}

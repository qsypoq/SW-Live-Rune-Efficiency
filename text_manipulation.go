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

func stat_exist(stat string) bool {
	for value, _ := range max_value {
		if stat == value {
			return true
		}
	}
	return false
}

func correct_stat(stat string) string {
	hit := 0
	try := 0
	if stat_exist(stat) {
		return stat
	} else {
		for value, _ := range max_value {
			hit = 0
			try = 0
			if len(stat) != len(value) {
				continue
			} else {
				for i := 0; i < len(value); i++ {
					if value[i] == stat[i] {
						hit = hit + 1
						try = try + 1
					} else {
						try = try + 1
					}
				}
				if float64(hit)/float64(try) > 0.66 {
					return value
				}

			}
		}
	}
	return "error"
}

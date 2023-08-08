package main

import (
	"regexp"
	"strings"
)

func trimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func split_stats(rune_subs string) []string {
	rune_substrimed := trimAllSpace(rune_subs)
	pattern_percent := "([%])"
	repercent := regexp.MustCompile(pattern_percent)
	percent_done := repercent.ReplaceAllString(rune_substrimed, "%|")
	sep := "|"
	pattern_digit := "([0-9]\\s+)"
	renumb := regexp.MustCompile(pattern_digit)
	indexes := renumb.FindAllStringIndex(percent_done, -1)
	move := 0
	for _, v := range indexes {
		p1 := v[0] + move
		p2 := v[1] + move
		percent_done = percent_done[:p1] + percent_done[p1:p2] + sep + percent_done[p2:]
		move += 2
	}
	result := strings.Split(percent_done, sep)
	result = result[:len(result)-1]
	return result
}

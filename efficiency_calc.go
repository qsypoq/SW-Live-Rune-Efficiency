package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var max_hit = map[string]int{
	"HPper":      8,
	"DEFper":     8,
	"ATKper":     8,
	"Accuracy":   8,
	"CRI Rate":   6,
	"CRI Dmg":    7,
	"ATK":        20,
	"DEF":        20,
	"HP":         375,
	"SPD":        6,
	"Resistance": 8,
}

func get_hit_value(value string, sub string) float64 {
	intvalue, _ := strconv.Atoi(value)
	hit := float64(intvalue) / float64(max_hit[sub])
	return hit
}

func get_hit_number(subs string, stat string) float64 {
	splitted := split_stats(subs)
	total_hits := float64(0)
	for _, v := range splitted {
		result := strings.Split(v, "+")
		subs := trimAllSpace(result[0])
		value := trimAllSpace(result[1])
		if strings.Contains(value, "%") {
			if strings.Contains(subs, "HP") || strings.Contains(subs, "ATK") || strings.Contains(subs, "DEF") {
				subs = subs + "per"
			}
			value = strings.Replace(value, "%", "", -1)
		}
		hit_number := get_hit_value(value, subs)
		total_hits = total_hits + hit_number
	}

	if strings.Count(stat, "+") == 2 {
		splitted := split_stats(stat)
		result := strings.Split(splitted[1], "+")
		subs := trimAllSpace(result[0])
		value := trimAllSpace(result[1])
		if strings.Contains(value, "%") {
			if strings.Contains(subs, "HP") || strings.Contains(subs, "ATK") || strings.Contains(subs, "DEF") {
				subs = subs + "per"
			}
			value = strings.Replace(value, "%", "", -1)
		}
		hit_inate := get_hit_value(value, subs)
		total_hits = total_hits + hit_inate
	}

	return total_hits
}

func compute_efficiency(hit_number float64) float64 {
	efficiency := 100 + ((hit_number - 9) * 11.11)
	efficiency = math.Round(efficiency*100) / 100
	return efficiency
}

func get_efficiency() (string, string, string, string) {
	generate_tmp_imgs()
	rune_name, rune_stats, rune_subs := get_rune_infos()
	current_efficiency := fmt.Sprintf("%.2f", compute_efficiency(get_hit_number(rune_subs, rune_stats)))
	return rune_name, rune_stats, rune_subs, current_efficiency
}

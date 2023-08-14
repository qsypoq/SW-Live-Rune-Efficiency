package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

var max_value = map[string]int{
	"HPper":      40,
	"DEFper":     40,
	"ATKper":     40,
	"Accuracy":   40,
	"CRI Rate":   30,
	"CRI Dmg":    35,
	"ATK":        100,
	"DEF":        100,
	"HP":         1875,
	"SPD":        30,
	"Resistance": 40,
}

func get_hit_value(value string, sub string) float64 {
	intvalue, _ := strconv.Atoi(value)
	hit := float64(intvalue) / float64(max_value[sub])
	return hit
}

func get_hit_number(subs string, stat string) float64 {
	splitted := split_stats(subs)
	total_hits := float64(0)

	for _, v := range splitted {
		if !strings.Contains(v, "+") {
			continue
		}
		result := strings.Split(v, "+")
		subs := trimAllSpace(result[0])
		subs = clean_char(subs, "[0-9]")
		value := trimAllSpace(result[1])

		if strings.Contains(value, "%") {
			if strings.Contains(subs, "HP") || strings.Contains(subs, "ATK") || strings.Contains(subs, "DEF") {
				subs = strings.TrimSpace(subs)
				subs = subs + "per"
			}
			value = strings.Replace(value, "%", "", -1)
		}
		hit_number := get_hit_value(value, subs)
		total_hits = total_hits + hit_number
		// fmt.Print(subs, ":", value, ":", hit_number, "\n")
	}

	if strings.Count(stat, "+") == 2 {
		stat = strings.Replace(stat, "©", "", -1)
		splitted := split_stats(stat)
		result := strings.Split(splitted[1], "+")
		subs := trimAllSpace(result[0])
		subs = clean_char(subs, "[0-9]")
		value := trimAllSpace(result[1])
		if strings.Contains(value, "%") {
			if strings.Contains(subs, "HP") || strings.Contains(subs, "ATK") || strings.Contains(subs, "DEF") {
				subs = strings.TrimSpace(subs)
				subs = subs + "per"
			}
			value = strings.Replace(value, "%", "", -1)
		}
		hit_inate := get_hit_value(value, subs)
		total_hits = total_hits + hit_inate
		// fmt.Print(subs, ":", value, ":", hit_inate, "\n")
	}

	return total_hits
}

func compute_efficiency(hit_number float64) float64 {
	efficiency := ((1 + hit_number) / 2.8) * 100
	return efficiency
}

func get_efficiency() (string, string, string, string) {
	rune_name, rune_stats, rune_subs := generate_rune()
	current_efficiency := fmt.Sprintf("%.2f", compute_efficiency(get_hit_number(rune_subs, rune_stats)))
	return clean_char(rune_name, "\n"), rune_stats, rune_subs, current_efficiency
}

func get_tier(efficiency string) (string, color.RGBA) {
	score, _ := strconv.ParseFloat(efficiency, 64)
	switch {
	case score < 85.7142857:
		return "Inate Rare Tier", color.RGBA{R: 67, G: 214, B: 215, A: 255}
	case score > 85.7142857 && score < 92.8571429:
		return "Inate Hero Tier", color.RGBA{R: 193, G: 17, B: 140, A: 255}
	case score > 92.8571429:
		return "Inate Legend Tier", color.RGBA{R: 187, G: 75, B: 28, A: 255}
	default:
		return "error", color.RGBA{0, 0, 0, 1}
	}
}

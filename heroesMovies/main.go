package main

import (
	"fmt"
	"strings"
)

func getDCMarvelHeroes(heroes []string) ([]string, []string) {
	marvel, dc := []string{}, []string{}

	for _, hero := range heroes {
		switch strings.ToLower(hero) {
		case "spiderman", "dr. strange":
			marvel = append(marvel, hero)
		case "batman", "superman":
			dc = append(dc, hero)
		}
	}

	return marvel, dc
}

func main() {
	heroes := []string{"Batman", "Spiderman", "Superman", "Dr. Strange"}
	marvel, dc := getDCMarvelHeroes(heroes)
	fmt.Println("Marvel Heroes: ", marvel)
	fmt.Println("DC Heroes: ", dc)
}

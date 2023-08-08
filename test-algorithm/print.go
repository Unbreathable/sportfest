package main

import "log"

func PrintYears(years []Year) {
	for _, year := range years {
		log.Println(year.Name, len(year.Player))
		for _, player := range year.Player {
			log.Println(player.Name)
		}
	}
}

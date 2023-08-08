package main

import (
	"log"
	"time"
)

var flawed = false
var flawCount = 0

func FoundFlaw(flaw string) {
	flawed = true
	flawCount = 0
}

func Compute(years []Year, games []Game, players []*Player, threads int) {

	for {
		// Run algorithm for computing players in each game
		SortAlgorithm(&years, &games, &players)

		// Create the teams
		team1, team2 := CreateTeams(&years, &players, &games)

		// Repair zero matches for players
		repairZeros(&years, &team1, &team2)

		// Fill empty matches
		fillEmptyMatches(&years, &team1, &team2)

		if flawed {
			log.Printf("Flaw found: %d\n", flawCount)
			log.Println("Resetting...")
			Reset(&years, &games, players)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Print debug stuff
		log.Println("Team 1:", len(team1))
		log.Println("Team 2:", len(team2))

		var matchCounter map[int][]*Player = map[int][]*Player{} // amount of games -> players
		for _, player := range players {
			matchCounter[len(player.Matches)] = append(matchCounter[len(player.Matches)], player)
		}

		log.Println("Empty matches:", GetEmptyMatches(&years))
		for _, players := range matchCounter {
			log.Println("Players with", len(players[0].Matches), "matches:", len(players))
		}

		var gameCounter map[string]int = map[string]int{}
		for _, year := range years {
			for _, game := range year.Playable {
				for _, match := range game.AvailableMatches {
					gameCounter[match.Game]++
				}
			}
		}

		for game, amount := range gameCounter {
			log.Println("Game", game, "has", amount, "matches")
		}

		break
	}
}

func Reset(years *[]Year, games *[]Game, players []*Player) {
	flawed = false
	flawCount = 0

	for i := range *years {
		year := (*years)[i]
		year.GameAmount = map[string]int{}
		year.Playable = map[string]*PlayableGame{}
	}

	for i := range players {
		player := players[i]
		player.Matches = []*Match{}
	}
}

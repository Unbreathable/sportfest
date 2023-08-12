package algorithm

import (
	"log"
)

var flawed = false
var flawCount = 0

func FoundFlaw(flaw string) {
	flawed = true
	flawCount = 0
}

func Compute(years []Year, games []Game, players []*Player, threads int) bool {

	// Run algorithm for computing players in each game
	SortAlgorithm(&years, &games, &players)

	// Create the teams
	team1, team2 := CreateTeams(&years, &players, &games)

	// Repair zero matches for players
	repairZeros(&years, &team1, &team2)

	// Fill empty matches
	fillEmptyMatches(&years, &team1, &team2)

	if flawed {
		log.Printf("Flaws found: %d\n", flawCount)
		return false
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

	var gameCounter map[uint]int = map[uint]int{}
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

	return true
}

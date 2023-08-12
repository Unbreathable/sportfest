package algorithm

import (
	"fmt"
	"log"
)

func Compute(years *[]Year, games *[]Game, players *[]*Player) ([]string, bool) {

	// Run algorithm for computing players in each game
	SortAlgorithm(years, games, players)

	// Create the teams
	team1, team2 := CreateTeams(years, players, games)

	// Repair zero matches for players
	repairZeros(years, &team1, &team2)

	// Fill empty matches
	fillEmptyMatches(years, &team1, &team2)

	// Print debug stuff
	log.Println("Team 1:", len(team1))
	log.Println("Team 2:", len(team2))

	var matchCounter map[int][]*Player = map[int][]*Player{} // amount of games -> players
	for _, player := range *players {
		matchCounter[len(player.Matches)] = append(matchCounter[len(player.Matches)], player)
	}

	logs := []string{}
	for i, players := range matchCounter {
		logs = append(logs, fmt.Sprintf("Players with %d matches: %d", i, len(players)))
	}
	logs = append(logs, fmt.Sprintf("Team 1: %d", len(team1)))
	logs = append(logs, fmt.Sprintf("Team 2: %d", len(team2)))

	return logs, true
}

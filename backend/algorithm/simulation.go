package algorithm

import (
	"log"

	"github.com/Unbreathable/sportfest/backend/util"
)

var Games = []Game{
	{0, 4, 6},
	{1, 6, 11},
	{2, 2, 2},
	{3, 5, 5},
	{4, 3, 7},
	{5, 4, 10},
	{6, 6, 10},
}

var Players = []*Player{}

func StartSimulation(playerAmount int, friendPairs int, playersPerYear int) {

	// Generate players
	for i := 0; i < playerAmount; i++ {
		Players = append(Players, &Player{
			ID:            uint(util.RandomInt(0, 10000000)),
			Friend:        0,
			Year:          0,
			GamesSelected: RandomGames(3, 5),
			Matches:       []*Match{},
		})
	}

	// Generate years
	yearAmount := playerAmount / playersPerYear
	years := []Year{}
	for i := 0; i < yearAmount; i++ {
		year := Year{
			ID:         uint(util.RandomInt(0, 10000000)),
			Player:     Players[i*playersPerYear : (i+1)*playersPerYear],
			GameAmount: map[uint]int{},
			Playable:   map[uint]*PlayableGame{},
		}
		years = append(years, year)

		// Assign year to players
		for _, player := range year.Player {
			player.Year = year.ID
		}
	}

	// Generate friend pairs
	for i := 0; i < friendPairs; i++ {

		// Get two random players
		player1 := Players[RandomInt(0, len(Players))]
		player2 := Players[RandomInt(0, len(Players))]

		// Check if they are already friends
		if player1.Friend == 0 && player2.Friend == 0 && player1.Year == player2.Year {
			player1.Friend = player2.ID
			player2.Friend = player1.ID
		}
	}

	// Run algorithm for computing players in each game
	logs, _ := Compute(&years, &Games, &Players)
	for _, l := range logs {
		log.Println(l)
	}
}

// Select a random amount of games between min and max
func RandomGames(min, max int) []uint {

	// Get a random amount of games
	amount := RandomInt(min, max)

	// Get a random selection of games
	games := []uint{}
	for i := 0; i < amount; i++ {

		// Check if games are already in the slice
		for {
			game := Games[RandomInt(0, len(Games))].ID
			if !contains(games, game) {
				games = append(games, game)
				break
			}
		}
	}

	return games
}

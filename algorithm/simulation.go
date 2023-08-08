package main

var Games = []Game{
	{"Völkerball", 4, 6},
	{"Fußball", 6, 11},
	{"Running", 2, 2},
	{"Basketball", 5, 5},
	{"Brennball", 3, 7},
	{"CTF", 4, 10},
	{"CTF2", 6, 10},
}

var Players = []*Player{}

func StartSimulation(playerAmount int, friendPairs int, playersPerYear int) {

	// Generate players
	for i := 0; i < playerAmount; i++ {
		Players = append(Players, &Player{
			Name:          GenerateToken(10),
			Friend:        "",
			Year:          "",
			GamesSelected: RandomGames(3, 5),
			Matches:       []*Match{},
		})
	}

	// Generate years
	yearAmount := playerAmount / playersPerYear
	years := []Year{}
	for i := 0; i < yearAmount; i++ {
		year := Year{
			Name:             GenerateToken(5),
			Player:           Players[i*playersPerYear : (i+1)*playersPerYear],
			GameAmount:       map[string]int{},
			Playable:         map[string]*PlayableGame{},
			PlayerSelections: map[string][]*Player{},
		}
		years = append(years, year)

		// Assign year to players
		for _, player := range year.Player {
			player.Year = year.Name
		}
	}

	// Generate friend pairs
	for i := 0; i < friendPairs; i++ {

		// Get two random players
		player1 := Players[RandomInt(0, len(Players))]
		player2 := Players[RandomInt(0, len(Players))]

		// Check if they are already friends
		if player1.Friend == "" && player2.Friend == "" && player1.Year == player2.Year {
			player1.Friend = player2.Name
			player2.Friend = player1.Name
		}
	}

	// Run algorithm for computing players in each game
	Compute(years, Games, Players, 1)
}

// Select a random amount of games between min and max
func RandomGames(min, max int) []string {

	// Get a random amount of games
	amount := RandomInt(min, max)

	// Get a random selection of games
	games := []string{}
	for i := 0; i < amount; i++ {

		// Check if games are already in the slice
		for {
			game := Games[RandomInt(0, len(Games))].Name
			if !contains(games, game) {
				games = append(games, game)
				break
			}
		}
	}

	return games
}

// Check if a string is present in a slice of strings
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

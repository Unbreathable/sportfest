package main

type Pair struct {
	PairSize         int
	Player1          *Player
	Player2          *Player
	OverlappingGames []string
}

func CreateTeams(years *[]Year, players *[]*Player, games *[]Game) (team1 []*Player, team2 []*Player) {

	// Generate map of games
	gameMap := map[string]*Game{}
	for i := range *games {
		game := (*games)[i]
		gameMap[game.Name] = &game
	}

	// Team blue and team red for the competition
	var team1Global []*Player = []*Player{}
	var team2Global []*Player = []*Player{}

	// Sort the players into teams
	for _, year := range *years {

		// Team blue and team red for the year
		var team1Local []*Player = []*Player{}
		var team2Local []*Player = []*Player{}

		// Group into pairs of friends
		var pairs []Pair = []Pair{}
		sorted := false
		for _, player := range year.Player {

			// Check if player's friend is already in a pair
			for _, pair := range pairs {
				if player.Friend == pair.Player1.Name {
					pair.Player2.Name = player.Name

					// Compute overlapping games
					for _, game := range pair.Player1.GamesSelected {
						if contains(pair.Player2.GamesSelected, game) {
							pair.PairSize = 2
							pair.OverlappingGames = append(pair.OverlappingGames, game)
						}
					}

					sorted = true
					break
				}
			}

			if sorted {
				continue
			}

			// Add to new pair
			pairs = append(pairs, Pair{Player1: player, Player2: &Player{}, PairSize: 1, OverlappingGames: player.GamesSelected})
			sorted = false
		}

		// Sort the pairs into teams of the matches
		for {
			if len(pairs) == 0 {
				break
			}

			// Get a random pair
			index := RandomInt(0, len(pairs))
			pair := pairs[index]

			// Compute best team for the pair based on overlapping games
			matches, team := ComputeBestTeamOverlapping(&pair, &year)

			// Add to matches
			addMatchesAndTeamToPlayer(pair.Player1, matches, team, &team1Local, &team2Local)
			if pair.PairSize == 2 {
				addMatchesAndTeamToPlayer(pair.Player2, matches, team, &team1Local, &team2Local)
			}

			pairs = append(pairs[:index], pairs[index+1:]...) // Remove from slice
		}

		// Add to global teams
		team1Global = append(team1Global, team1Local...)
		team2Global = append(team2Global, team2Local...)
	}

	return team1Global, team2Global
}

func addMatchesAndTeamToPlayer(player *Player, matches []*Match, team int, team1 *[]*Player, team2 *[]*Player) {

	// Add to matches teams
	for _, match := range matches {

		if team == 0 {
			match.Team1 = append(match.Team1, player)
		} else {
			match.Team2 = append(match.Team2, player)
		}
	}

	// Add to team
	if team == 0 {
		*team1 = append(*team1, player)
	} else {
		*team2 = append(*team2, player)
	}

	// Add to player
	player.Matches = append(player.Matches, matches...)
}

func ComputeBestTeamOverlapping(pair *Pair, year *Year) ([]*Match, int) {

	// Compute best team for the pair
	var team1Matches = []*Match{}
	var team2Matches = []*Match{}
	for _, game := range pair.OverlappingGames {

		if year.Playable[game] == nil {
			continue
		}

		matches := year.Playable[game].AvailableMatches
		if len(matches) == 0 {
			continue
		}

		// Get perfect team in the match
		for _, match := range matches {

			// Check if match is full
			spaceTeam1 := match.TeamSize - len(match.Team1)
			spaceTeam2 := match.TeamSize - len(match.Team2)
			if spaceTeam1 < pair.PairSize && spaceTeam2 < pair.PairSize {
				continue
			}

			// Compute best team for the pair
			if spaceTeam1 >= spaceTeam2 && spaceTeam1 >= pair.PairSize {
				team1Matches = append(team1Matches, match)
			} else if spaceTeam2 >= spaceTeam1 && spaceTeam2 >= pair.PairSize {
				team2Matches = append(team2Matches, match)
			}
		}
	}

	// Check if there are any matches
	if len(team1Matches) == 0 && len(team2Matches) == 0 {
		FoundFlaw("No matches found for pair: " + pair.Player1.Name + " + " + pair.Player2.Name)
	}

	// Eliminate matches with same game
	team1Matches = removeDuplicates(team1Matches)
	team2Matches = removeDuplicates(team2Matches)

	// Get better team
	if len(team1Matches) > len(team2Matches) {
		return team1Matches, 0
	} else {
		return team2Matches, 1
	}
}

func removeDuplicates(matches []*Match) []*Match {
	var filtered []*Match = []*Match{}

	// Remove duplicates
	for _, match := range matches {
		found := false
		for _, matchFiltered := range filtered {
			if match.Game == matchFiltered.Game {
				found = true
				break
			}
		}

		if found {
			continue
		}

		filtered = append(filtered, match)
	}

	return filtered
}

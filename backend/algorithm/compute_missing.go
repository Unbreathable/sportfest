package algorithm

func GetEmptyMatches(years *[]Year) int {

	var emptyMatches []*Match = []*Match{}

	for _, year := range *years {
		for _, game := range year.Playable {
			for _, match := range game.AvailableMatches {
				if len(match.Team1) < game.MinTeamSize || len(match.Team2) < game.MinTeamSize {
					if len(game.AvailableMatches) == 1 {
						emptyMatches = append(emptyMatches, match)
					}
				}
			}
		}
	}

	return len(emptyMatches)
}

func fillEmptyMatches(years *[]Year, team1 *[]*Player, team2 *[]*Player) {

	for _, year := range *years {
		var team1ToFill []*Match = []*Match{}
		var team2ToFill []*Match = []*Match{}

		var team1MatchCounter map[int][]*Player = map[int][]*Player{} // amount of games -> players
		var team2MatchCounter map[int][]*Player = map[int][]*Player{} // amount of games -> players

		// Compute match counters
		for _, player := range *team1 {
			if player.Year == year.ID {
				team1MatchCounter[len(player.Matches)] = append(team1MatchCounter[len(player.Matches)], player)
			}
		}

		for _, player := range *team2 {
			if player.Year == year.ID {
				team2MatchCounter[len(player.Matches)] = append(team2MatchCounter[len(player.Matches)], player)
			}
		}

		// Compute matches to fill
		for _, game := range year.Playable {
			for _, match := range game.AvailableMatches {

				// Check if not full
				if len(match.Team1) != match.TeamSize {
					team1ToFill = append(team1ToFill, match)
				}

				if len(match.Team2) != match.TeamSize {
					team2ToFill = append(team2ToFill, match)
				}
			}
		}

		// Fill matches
		FillTeam(team1, &team1ToFill, &team1MatchCounter)
		FillTeam(team2, &team2ToFill, &team2MatchCounter)
	}
}

func FillTeam(team *[]*Player, matches *[]*Match, matchCounter *map[int][]*Player) {
	for _, match := range *matches {
		for i := 0; i < match.TeamSize-len(match.Team1); i++ {

			// Get player with least matches
			var player *Player = nil
			playerIndex := 0
			for _, players := range *matchCounter {

				if player != nil {
					break
				}

				if len(players) > 0 {
					for index, p := range players {

						// Check if player is already in a match
						found := false
						for _, m := range p.Matches {
							if m.Game == match.Game {
								found = true
								continue
							}
						}

						if found {
							continue
						}

						// Check if player is interested
						if contains(p.GamesSelected, match.Game) {
							player = p
							playerIndex = index
						}
						break
					}
				}
			}

			if player == nil {
				continue
			}

			// Remove from previous match counter
			(*matchCounter)[len(player.Matches)] = append((*matchCounter)[len(player.Matches)][:playerIndex], (*matchCounter)[len(player.Matches)][playerIndex+1:]...)

			// Add to team
			match.Team1 = append(match.Team1, player)
			player.Matches = append(player.Matches, match)
			(*matchCounter)[len(player.Matches)] = append((*matchCounter)[len(player.Matches)], player)
		}
	}
}

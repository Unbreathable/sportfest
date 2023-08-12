package algorithm

func repairZeros(years *[]Year, team1 *[]*Player, team2 *[]*Player) {

	// Repair zero matches for players
	for _, year := range *years {
		for _, player := range year.Player {
			if len(player.Matches) == 0 {

				// Get best team
				matches, team := computeBestTeam(player, &year)

				// Add to matches teams
				addMatchesAndTeamToPlayer(player, matches, team, team1, team2)
			}
		}
	}

}

func computeBestTeam(player *Player, year *Year) ([]*Match, int) {

	// Compute best team for the pair
	var team1Matches = []*Match{}
	var team2Matches = []*Match{}
	for _, game := range player.GamesSelected {

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
			if spaceTeam1 == 0 && spaceTeam2 == 0 {
				continue
			}

			// Compute best team for the pair
			if spaceTeam1 >= spaceTeam2 && spaceTeam1 > 0 {
				team1Matches = append(team1Matches, match)
			} else if spaceTeam2 >= spaceTeam1 && spaceTeam2 > 0 {
				team2Matches = append(team2Matches, match)
			}
		}
	}

	// Check if there are any matches
	if len(team1Matches) == 0 && len(team2Matches) == 0 {
		FoundFlaw("No matches found for player")
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

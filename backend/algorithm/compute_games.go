package algorithm

func SortAlgorithm(years *[]Year, games *[]Game, players *[]*Player) {

	// Generate map of games
	gameMap := map[uint]Game{}
	for i := range *games {
		game := (*games)[i]
		gameMap[game.ID] = game
	}

	// Get the number of players in each game per year
	for _, game := range *games {
		for _, year := range *years {
			for _, player := range year.Player {
				if contains(player.GamesSelected, game.ID) {
					year.GameAmount[game.ID] += 1
				}
			}
		}
	}

	// Print the number of players in each game per year
	for _, year := range *years {
		for game, amount := range year.GameAmount {

			// Check if it can be played
			if amount > gameMap[game].MinTeamSize*2 {
				playableGame := PlayableGame{
					Game:             gameMap[game],
					AvailableMatches: []*Match{},
					LeftPlayerAmount: 0,
				}

				// Compute
				matches := amount / (gameMap[game].MinTeamSize * 2)
				playableGame.LeftPlayerAmount += amount % (gameMap[game].MinTeamSize * 2)
				for i := 0; i < matches; i++ {
					playableGame.AvailableMatches = append(playableGame.AvailableMatches, &Match{
						TeamSize: gameMap[game].MinTeamSize,
						Game:     game,
					})
				}

				// Try and fill up the teams with left-over players
				if playableGame.LeftPlayerAmount != 0 {
					restPairs := playableGame.LeftPlayerAmount / 2
					if restPairs == 0 {
						continue
					}

					for {
						if restPairs == 0 {
							break
						}

						// Add remaining players to teams with space
						space := false
						for i, match := range playableGame.AvailableMatches {
							if match.TeamSize < gameMap[game].MaxTeamSize {
								space = true
								playableGame.AvailableMatches[i].TeamSize += 1
								restPairs--
								if restPairs == 0 {
									break
								}
							}
						}

						if !space {
							break
						}
					}
				}

				// Add to playable games
				year.Playable[game] = &playableGame
			}
		}
	}
}

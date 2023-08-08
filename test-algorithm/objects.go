package main

type Game struct {
	Name        string
	MinTeamSize int
	MaxTeamSize int
}

type PlayableGame struct {
	Game
	AvailableMatches []*Match
	LeftPlayerAmount int
}

type Match struct {
	ID       string
	Game     string
	TeamSize int
	Team1    []*Player
	Team2    []*Player
}

func (m *Match) GetTeam(team int) []*Player {
	if team == 0 {
		return m.Team1
	} else {
		return m.Team2
	}
}

type EzYear struct {
	Name             string
	Player           []Player
	GameAmount       map[string]int
	Playable         map[string]*PlayableGame
	PlayerSelections map[string][]*Player
}

type Year struct {
	Name             string
	Player           []*Player
	GameAmount       map[string]int
	Playable         map[string]*PlayableGame
	PlayerSelections map[string][]*Player
}

type Player struct {
	Name          string
	Friend        string
	Year          string
	GamesSelected []string // slice of game names
	Matches       []*Match
}

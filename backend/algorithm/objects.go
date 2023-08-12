package algorithm

type Game struct {
	ID          uint
	MinTeamSize int
	MaxTeamSize int
}

type PlayableGame struct {
	Game
	AvailableMatches []*Match
	LeftPlayerAmount int
}

type Match struct {
	Game     uint
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

type Year struct {
	ID         uint
	Player     []*Player
	GameAmount map[uint]int
	Playable   map[uint]*PlayableGame
}

type Player struct {
	ID            uint
	Friend        uint
	Year          uint
	GamesSelected []uint // slice of game ids
	Matches       []*Match
}

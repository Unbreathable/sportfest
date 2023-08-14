package database

const Team1 = 0
const Team2 = 1

type Game struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Video       string `json:"video" gorm:"not null"`
	MinTeamSize uint   `json:"minTeamSize" gorm:"not null"`
	MaxTeamSize uint   `json:"maxTeamSize" gorm:"not null"`
}

type Match struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	Year     uint `json:"year" gorm:"not null"` // Year ID
	Game     uint `json:"game" gorm:"not null"`
	TeamSize int  `json:"teamSize" gorm:"not null"`
}

type TeamMember struct {
	ID    uint `json:"id" gorm:"primaryKey"`
	Team  uint `json:"team" gorm:"not null"` // Team1 or Team2 (0 or 1)
	Match uint `json:"match" gorm:"not null"`
	Game  uint `json:"game" gorm:"not null"`
	User  uint `json:"user" gorm:"not null"`
}

type Year struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type AdminAccount struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	Code string `json:"code" gorm:"not null"`
	Team uint   `json:"team"`                 // Team1 or Team2 (0 or 1)
	Year uint   `json:"year" gorm:"not null"` // Year ID
}

type Choice struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	User   uint `json:"user" gorm:"not null"`   // User ID
	Choice uint `json:"choice" gorm:"not null"` // Game choice
}

type Friendship struct {
	ID         uint   `json:"id" gorm:"primaryKey"` // Account ID
	Code       string `json:"code" gorm:"not null"` // Friend code
	FriendCode string `json:"friendCode"`           // The friend code of the other user
}

const FlawTypeTeamInvalid = "team_invalid"
const FlawTypeTeamEmpty = "team_empty"
const FlawTypePlayerNotHandled = "player_not_handled"

// Flaw is a mistake found by the flaw algorithm
type Flaw struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Type        string `json:"type" gorm:"not null"`        // Type of flaw
	Description string `json:"description" gorm:"not null"` // Description of flaw
	Target      uint   `json:"target" gorm:"not null"`      // User ID
}

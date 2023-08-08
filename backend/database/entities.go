package database

const Team1 = 0
const Team2 = 1

type Game struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	MinTeamSize int    `json:"minTeamSize"`
	MaxTeamSize int    `json:"maxTeamSize"`
}

type Match struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	Game     uint `json:"game" gorm:"not null"`
	TeamSize int  `json:"teamSize" gorm:"not null"`
}

type TeamMember struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	Team    uint `json:"team" gorm:"not null"` // Team1 or Team2 (0 or 1)
	Match   uint `json:"match" gorm:"not null"`
	Account uint `json:"name" gorm:"not null"`
}

type Year struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"not null"`
	Mapping string `json:"mapping" gorm:"not null"` // In case we need to import a csv or something
}

type Account struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Year     uint   `json:"year" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Admin    bool   `json:"admin" gorm:"not null"`
}

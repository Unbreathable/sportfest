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
	ID      uint `json:"id" gorm:"primaryKey"`
	Team    uint `json:"team" gorm:"not null"` // Team1 or Team2 (0 or 1)
	Match   uint `json:"match" gorm:"not null"`
	Game    uint `json:"game" gorm:"not null"`
	Account uint `json:"name" gorm:"not null"`
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
	Year uint   `json:"year" gorm:"not null"` // Year ID
}

type Choice struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	User   uint `json:"user" gorm:"not null"`   // User ID
	Choice uint `json:"choice" gorm:"not null"` // Game choice
}

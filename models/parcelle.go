package models

import "github.com/google/uuid"

type Parcelle struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Libelle         string    `json:"libelle"`
	Geolocalisation string    `json:"geolocalisation"`
	Surface         float64   `json:"surface"`
	UserID          uuid.UUID `json:"user_id" gorm:"type:uuid"`
}

func (Parcelle) TableName() string {
	return "parcelle"
}

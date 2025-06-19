package models

import "github.com/google/uuid"

type TypeCulture struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Libelle       string    `json:"libelle"`
	PrixBordChamp float64   `json:"prix_bord_champ"`
}

func (TypeCulture) TableName() string {
	return "type_culture"
}

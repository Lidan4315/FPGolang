package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Merek struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Nama string    `json:"nama"`
}

type Mobil struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	MerekID         uuid.UUID `gorm:"type:uuid" json:"merek_id"`
	Merek           Merek     `gorm:"foreignKey:MerekID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Type            string    `json:"type"`
	NoPlat          string    `json:"no_plat"`
	Warna           string    `json:"warna"`
	InitialCondition string   `json:"initial_condition"`
	Harga           float64   `json:"harga"`
	Deskripsi       string    `json:"deskripsi"`
	ImageUrl        string    `json:"image_url"`

	Timestamp
}

func (m *Mobil) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return nil
}

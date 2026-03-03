package models

import (
	"time"

	"github.com/google/uuid"
)

type Lomba struct {
	InternalID int64 `json:"-" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID uuid.UUID `json:"public_id" db:"public_id"`
	NamaLomba string `json:"nama_lomba" db:"nama_lomba"`
	DeskripsiLomba string `json:"deskripsi_lomba" db:"deskripsi_lomba"`
	Persyaratan string `json:"persyaratan" db:"persyaratan"`
	TglBuka time.Time `json:"tgl_buka" db:"tgl_buka"`
	TglTutup time.Time `json:"tgl_tutup" db:"tgl_tutup"`
	
	// Relasi
	CreatedBy int64 `json:"created_by" db:"created_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (Lomba) TableName() string {
    return "lomba"
}
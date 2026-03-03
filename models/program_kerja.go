package models

import (
	"time"

	"github.com/google/uuid"
)

type ProgramKerja struct {
	InternalID int64 `json:"-" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID uuid.UUID `json:"public_id" db:"public_id"`
	NamaProker string `json:"nama_proker" db:"nama_proker"`
	Deskripsi string `json:"deskripsi" db:"deskripsi"`
	Divisi string `json:"divisi" db:"divisi"`
	Status string `json:"status" db:"status"`
	LinkOprec string `json:"link_oprec" db:"link_oprec"`

	// Relasi 
	CreatedBy int64 `json:"created_by" db:"created_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (ProgramKerja) TableName() string {
    return "program_kerja"
}
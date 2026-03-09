package models

import (
	"time"

	"github.com/google/uuid"
)

type Beasiswa struct {
	InternalID int64 `json:"-" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID uuid.UUID `json:"public_id" db:"public_id"`
	NamaBeasiswa string `json:"nama_beasiswa" db:"nama_beasiswa"`
	LinkPendaftaran string `json:"link_pendaftaran" db:"link_pendaftaran"`
	TglBuka time.Time `json:"tgl_buka" db:"tgl_buka"`
	TglTutup time.Time `json:"tgl_tutup" db:"tgl_tutup"`

	// Relasi
	CreatedBy int64 `json:"created_by" db:"created_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (Beasiswa) TableName() string {
    return "beasiswa"
}

type BeasiswaResp struct {
	PublicID uuid.UUID `json:"public_id" db:"public_id"`
	NamaBeasiswa string `json:"nama_beasiswa" db:"nama_beasiswa"`
	LinkPendaftaran string `json:"link_pendaftaran" db:"link_pendaftaran"`
	TglBuka time.Time `json:"tgl_buka" db:"tgl_buka"`
	TglTutup time.Time `json:"tgl_tutup" db:"tgl_tutup"`
}
package models

import (
	"time"
	"gorm.io/gorm"
)

type MediaType struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	TypeName  string         `json:"type_name" gorm:"unique;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Media struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	MediaTypeID uint           `json:"media_type_id" gorm:"not null"`
	MediaType   MediaType      `json:"media_type" gorm:"foreignKey:MediaTypeID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type MediaTypeFormat struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Format      string         `json:"format"`
	MediaTypeID uint           `json:"media_type_id" gorm:"not null"`
	MediaType   MediaType      `json:"media_type" gorm:"foreignKey:MediaTypeID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type ListeType struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	TypeName  string         `json:"type_name" gorm:"unique;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Liste struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Navn        string         `json:"navn"`
	ListeTypeID uint           `json:"liste_type_id" gorm:"not null"`
	ListeType   ListeType      `json:"liste_type" gorm:"foreignKey:ListeTypeID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	LagetAv     string         `json:"laget_av"`
}

type ListeInstans struct {
	ID                                  uint           `json:"id" gorm:"primaryKey"`
	ListeID                            uint           `json:"liste_id" gorm:"not null"`
	Liste                              Liste          `json:"liste" gorm:"foreignKey:ListeID"`
	CreatedAt                          time.Time      `json:"created_at"`
	UpdatedAt                          time.Time      `json:"updated_at"`
	DeletedAt                          gorm.DeletedAt `json:"-" gorm:"index"`
	AntallDagerPerTick                 int            `json:"antall_dager_per_tick"`
	AntallTillatteAnmeldelserPerMedia  int            `json:"antall_tillatte_anmeldelser_per_media"`
	AntallDagerMellomAnmeldelser       int            `json:"antall_dager_mellom_anmeldelser"`
	Ferdig                             bool           `json:"ferdig"`
}

type ListeInstansTick struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	ListeInstansID  uint           `json:"liste_instans_id" gorm:"not null"`
	ListeInstans    ListeInstans   `json:"liste_instans" gorm:"foreignKey:ListeInstansID"`
	MediaID         uint           `json:"media_id" gorm:"not null"`
	Media           Media          `json:"media" gorm:"foreignKey:MediaID"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

type ListeMediaTilhorighet struct {
	MediaID uint  `json:"media_id" gorm:"primaryKey"`
	Media   Media `json:"media" gorm:"foreignKey:MediaID"`
	ListeID uint  `json:"liste_id" gorm:"primaryKey"`
	Liste   Liste `json:"liste" gorm:"foreignKey:ListeID"`
}

type Bruker struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Brukernavn  string         `json:"brukernavn"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Gruppe struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Navn      string         `json:"navn"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type GruppeTilhorighet struct {
	BrukerID uint   `json:"bruker_id" gorm:"primaryKey"`
	Bruker   Bruker `json:"bruker" gorm:"foreignKey:BrukerID"`
	GruppeID uint   `json:"gruppe_id" gorm:"primaryKey"`
	Gruppe   Gruppe `json:"gruppe" gorm:"foreignKey:GruppeID"`
}

type BrukerListeInstansTilhorighet struct {
	BrukerID       uint         `json:"bruker_id" gorm:"primaryKey"`
	Bruker         Bruker       `json:"bruker" gorm:"foreignKey:BrukerID"`
	ListeInstansID uint         `json:"liste_instans_id" gorm:"primaryKey"`
	ListeInstans   ListeInstans `json:"liste_instans" gorm:"foreignKey:ListeInstansID"`
}

type GruppeListeInstansTilhorighet struct {
	GruppeID       uint         `json:"gruppe_id" gorm:"primaryKey"`
	Gruppe         Gruppe       `json:"gruppe" gorm:"foreignKey:GruppeID"`
	ListeInstansID uint         `json:"liste_instans_id" gorm:"primaryKey"`
	ListeInstans   ListeInstans `json:"liste_instans" gorm:"foreignKey:ListeInstansID"`
}

type MetadataSchema struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	MediaTypeID uint           `json:"media_type_id" gorm:"not null"`
	MediaType   MediaType      `json:"media_type" gorm:"foreignKey:MediaTypeID"`
	KeyName     string         `json:"key_name" gorm:"not null"`
	ValueType   string         `json:"value_type" gorm:"not null"`
	IsRequired  bool           `json:"is_required" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type MediaMetadata struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	MediaID  uint           `json:"media_id" gorm:"not null;uniqueIndex:idx_media_schema"`
	Media    Media          `json:"media" gorm:"foreignKey:MediaID"`
	SchemaID uint           `json:"schema_id" gorm:"not null;uniqueIndex:idx_media_schema"`
	Schema   MetadataSchema `json:"schema" gorm:"foreignKey:SchemaID"`
	Value    string         `json:"value" gorm:"not null"`
}

type Anmeldelse struct {
	ID                  uint               `json:"id" gorm:"primaryKey"`
	AntallStjerner      int                `json:"antall_stjerner"`
	Kommentar           string             `json:"kommentar"`
	MediaID             uint               `json:"media_id" gorm:"not null"`
	Media               Media              `json:"media" gorm:"foreignKey:MediaID"`
	MediaTypeFormatID   *uint              `json:"media_type_format_id"`
	MediaTypeFormat     *MediaTypeFormat   `json:"media_type_format" gorm:"foreignKey:MediaTypeFormatID"`
	BrukerID            uint               `json:"bruker_id" gorm:"not null"`
	Bruker              Bruker             `json:"bruker" gorm:"foreignKey:BrukerID"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
	DeletedAt           gorm.DeletedAt     `json:"-" gorm:"index"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Album struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

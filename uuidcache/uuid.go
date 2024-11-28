package uuidcache

import "time"

const (
	CacheEntryTable = "uuid_cache_uuids"
	NamesTable      = "uuid_cache_names"
)

type UuidCache interface {
	GetUUID(name string) (string, error)
	GetLatestName(uuid string) (string, error)
	GetNames(uuid string) ([]Name, error)
	Put(name, uuid string) error
}

type Name struct {
	Uuid         string    `gorm:"column:uuid;primaryKey;"`
	Name         string    `gorm:"column:name"`
	OriginalName string    `gorm:"column:original_name"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Name) TableName() string {
	return "uuid_cache_names"
}

type CacheEntry struct {
	Uuid      string    `gorm:"column:uuid;primaryKey;unique" json:"uuid"`
	Names     []Name    `gorm:"foreignKey:uuid" json:"names"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"omit"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (CacheEntry) TableName() string {
	return "uuid_cache_uuids"
}

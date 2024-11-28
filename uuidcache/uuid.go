package uuidcache

import "time"

type UuidCache interface {
	GetUUID(name string) (string, error)
	GetLatestName(uuid string) (string, error)
	GetNames(uuid string) ([]string, error)
	Put(name, uuid string) error
}

type Name struct {
	Uuid      string    `gorm:primaryKey,unique`
	Name      string    `gorm:""`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Name) TableName() string {
	return "uuid_cache_names"
}

type CacheEntry struct {
	Uuid      string    `gorm:"primaryKey,unique" json:"uuid"`
	Names     []Name    `gorm:"foreignKey:Uuid" json:"names"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"omit"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (CacheEntry) TableName() string {
	return "uuid_cache_uuids"
}

package uuidcache

import "gorm.io/gorm"

type SqlUuidCache struct {
	db *gorm.DB
}

func (s *SqlUuidCache) GetUUID(name string) (string, error) {
	return "", nil
}

func (s *SqlUuidCache) GetLatestName(uuid string) (string, error) {
	return "", nil
}

func (s *SqlUuidCache) GetNames(uuid string) ([]string, error) {
	return []string{}, nil
}

func (s *SqlUuidCache) Put(name, uuid string) error {
	return nil
}

func NewSqlUuidCache(db *gorm.DB) UuidCache {
	return &SqlUuidCache{
		db: db,
	}
}

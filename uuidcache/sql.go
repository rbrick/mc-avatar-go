package uuidcache

import (
	"strings"

	"gorm.io/gorm"
)

type SqlUuidCache struct {
	db *gorm.DB
}

func (s *SqlUuidCache) GetUUID(name string) (string, error) {
	var entry Name
	err := s.db.Table(NamesTable).Select("name", "uuid", "created_at").
		Where("name = ?", strings.ToLower(name)).Order("created_at DESC").
		First(&entry).Error

	if err != nil {
		return "", err
	}

	return entry.Uuid, nil
}

func (s *SqlUuidCache) GetLatestName(uuid string) (string, error) {
	var entry Name
	err := s.db.Table(NamesTable).Select("original_name", "uuid", "created_at").
		Where("uuid = ?", uuid).Order("created_at DESC").
		First(&entry).Error

	if err != nil {
		return "", err
	}

	return entry.OriginalName, nil
}

func (s *SqlUuidCache) GetNames(uuid string) ([]Name, error) {
	var entries []Name
	err := s.db.Table(NamesTable).Select("original_name", "uuid", "created_at").
		Where("uuid = ?", uuid).Order("created_at DESC").
		Find(&entries).Error

	if err != nil {
		return []Name{}, err
	}

	return entries, nil
}

func (s *SqlUuidCache) Put(name, uuid string) error {
	var entryCount int64

	s.db.Table(CacheEntryTable).
		Where("uuid = ?", uuid).
		Count(&entryCount)

	if entryCount == 0 {
		s.db.Table(CacheEntryTable).Create(&CacheEntry{
			Uuid: uuid,
		})
	}

	s.db.Table(NamesTable).Where("uuid = ?", uuid).
		Where("name = ?", strings.ToLower(name)).
		Count(&entryCount)

	if entryCount == 0 {
		s.db.Table(NamesTable).Create(&Name{
			Uuid:         uuid,
			Name:         strings.ToLower(name),
			OriginalName: name,
		})
	}

	return nil
}

func NewSqlUuidCache(db *gorm.DB) UuidCache {
	return &SqlUuidCache{
		db: db,
	}
}

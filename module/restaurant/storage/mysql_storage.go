package storage

import "gorm.io/gorm"

type mysqlStorage struct {
	DB *gorm.DB
}

func NewMysqlStorage(DB *gorm.DB) *mysqlStorage {
	return &mysqlStorage{DB: DB}
}

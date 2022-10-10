package file

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Uri         string
	Tags        []Tag
}

type Tag struct {
	gorm.Model
	Key    string `binding:"required"`
	FileID uint
}

type MetaData struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Tags        []Tag
}

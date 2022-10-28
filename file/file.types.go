package file

import (
	"time"
)

type File struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Uri         string
	Tags        []Tag
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tag struct {
	Key string `binding:"required"`
}

type MetaData struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Tags        []Tag
}

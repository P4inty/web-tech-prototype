package migration

import (
	"webtech/prototype/db"
	"webtech/prototype/file"
)

func Migrate() {
	db.DB.AutoMigrate(&file.File{}, &file.Tag{})
}

package databases

import "waservice/app/models"

func Migrate() {
	Db.AutoMigrate(
		&models.User{},
	)
}

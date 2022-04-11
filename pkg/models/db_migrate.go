package models

import "gorm.io/gorm"

var modelList = []DBModel{
	&User{},
	&Comp{},
	&Job{},
}

func migrate(tx *gorm.DB) error {

	migrator := tx.Migrator()
	for i := range modelList {
		if err := migrator.AutoMigrate(modelList[i]); err != nil {
			return err
		}
	}
	return nil
}

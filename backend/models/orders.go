package models

import "gorm.io/gorm"

type Orders struct {
	Id         uint `gorm:"primary key;autoIncrement" json:"id"`
	Product_id int  `json:"product_id"`
	Quantity   int  `json:"quantity"`
	Owner      int  `json:"owner"`
}

func MigrateOrders(db *gorm.DB) error {
	err := db.AutoMigrate(&Orders{})
	return err
}

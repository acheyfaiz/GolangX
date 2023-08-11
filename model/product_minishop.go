package model

import "time"

type ProductInMinishops struct {
	ID           int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Pid          int       `json:"pid" gorm:"column:pid"`
	Productid    int       `json:"productid" gorm:"column:productid"`
	Dateapply    time.Time `gorm:"column:dateapply"`
	Status       int       `json:"status" gorm:"column:status"`
	Dateapproved time.Time `gorm:"column:dateapply"`
	Itemsold     int       `json:"itemsold" gorm:"column:itemsold"`
}

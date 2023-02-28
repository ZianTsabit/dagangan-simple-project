package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	id		  int 		`json:"id" 			gorm:"primaryKey; not null; autoIncrement; unique; type:int;"`
	name      string 	`json:"name"		gorm:"unique; type:varchar(100)"`
	price     float64 	`json:"price"		gorm:"type:decimal(10,2)"`
	quantity  int 		`json:"quantity"	gorm:"type:int"`
}
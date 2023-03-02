package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id	  		int 		`json:"id" 			gorm:"primaryKey; not null; autoIncrement; unique; type:int;"`
	Name      	string 		`json:"name"		gorm:"unique; type:varchar(100)"`
	Price    	float64 	`json:"price"		gorm:"type:decimal(10,2)"`
	Quantity  	int 		`json:"quantity"	gorm:"type:int"`
}
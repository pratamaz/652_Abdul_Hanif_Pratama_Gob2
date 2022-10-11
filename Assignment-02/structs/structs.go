package structs

import (
	"time"
)

type Order struct {
	OrderId      int        `gorm:"primaryKey;autoIncrement;"`
	CustomerName string		
	Items 		 []Item	 	
	OrderedAt      time.Time 
}

type Item struct {
	ItemId   	int			`gorm:"primaryKey;autoIncrement;"` 
	ItemCode 	string		 
	Desription 	string		 
	Quantity	int			 
	OrderId		int			 `gorm:"references:OrderId"`
}
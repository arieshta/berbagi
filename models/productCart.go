package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCart struct {
	// Donor here means someone who is doing the act of giving, not as in donor roles
	// we limit donor to user with donor role and volunteer role
	// we limit recipient to children
	DonorID   			uint           	`gorm:"primaryKey" json:"donor_id"`
	RecipientID			uint           	`gorm:"primaryKey" json:"recipient_id" form:"recipient_id"`
	ProductPackageID 	uint           	`gorm:"not null;primaryKey" json:"product_package_id" form:"product_package_id"`
	Quantity  			int            	`gorm:"not null" json:"quantity" form:"quantity"`
	CreatedAt 			time.Time      	`json:"-"`
	UpdatedAt 			time.Time      	`json:"-"`
	DeletedAt 			gorm.DeletedAt 	`gorm:"index"`
	User     			User          	`gorm:"foreignKey:DonorID"`
	Children 			Children 		`gorm:"foreignKey:RecipientID"`
	ProductPackage   	ProductPackage  `gorm:"foreignKey:ProductPackageID"`
}

type ProductCartDelAPI struct {
	RecipientID			uint           	`json:"recipient_id" form:"recipient_id"`
	ProductPackageID 	uint           	`json:"product_package_id" form:"product_package_id"`
}	

type ProductCartGetAPI struct {
	RecipientID			uint           	`json:"recipient_id" form:"recipient_id"`
	ProductPackageID 	uint           	`json:"product_package_id" form:"product_package_id"`
	Quantity  			int            	`json:"quantity" form:"quantity"`
	Name 				string 			`json:"product_name"`
	Price  				int  			`json:"price"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

// Karena login behavior bergantung pada value role query, login controller harus mengekstrak informasi tersebut
// dan melakukan binding ke satu struktur yang uniform seperti dibawah ini
type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(75)" json:"name"`
	RoleID 	  uint		 	 `json:"role_id"`
	NIK       string         `gorm:"unique; type:varchar(16)" json:"nik"`
	Email     string         `gorm:"unique; type:varchar(50)" json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type LoginUserAPI struct {
	Email    	string	`json:"email"`
	Password 	string	`json:"password"`
}

type LoginSearchAPI struct {
	ID      	uint    `json:id`
	Email    	string	`json:"email"`
	Password 	string	`json:"password"`
	RoleID     	uint  	`json:"role_id"`
	RoleName 	string  `json:"role_name"`
}

type LoginResponseAPI struct {
	Message string `json:"message"`
	Token   string	`json:"token"`
}

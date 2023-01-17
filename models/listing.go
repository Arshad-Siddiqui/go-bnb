package models

import (
	"gorm.io/gorm"
)

type Listing struct {
	gorm.Model
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	// HostID      uint   `json:"host_id"`
}

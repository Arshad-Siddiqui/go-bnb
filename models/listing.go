package models

import (
	"gorm.io/gorm"
)

type Listing struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	// HostID      uint   `json:"host_id"`
}

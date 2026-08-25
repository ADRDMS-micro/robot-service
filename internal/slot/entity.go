package slot

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SlotStatus string

const (
	SlotEmpty      SlotStatus = "EMPTY"
	SlotReserved   SlotStatus = "RESERVED"
	SlotLoaded     SlotStatus = "LOADED"
	SlotDelivering SlotStatus = "DELIVERING"
)

type RobotSlot struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	RobotID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	SlotNumber int            `gorm:"type:integer;not null"`
	Status     SlotStatus     `gorm:"type:varchar(20);not null;default:'EMPTY'"`
	OrderID    *uuid.UUID     `gorm:"type:uuid"`
	StoreID    *uuid.UUID     `gorm:"type:uuid"`
	ReservedAt *time.Time     `gorm:""`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

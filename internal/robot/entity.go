package robot

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RobotStatus string

const (
	StatusIdle        RobotStatus = "IDLE"
	StatusEnRoute     RobotStatus = "EN_ROUTE"
	StatusMaintenance RobotStatus = "MAINTENANCE"
	StatusReturning   RobotStatus = "RETURNING"
)

type Robot struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string         `gorm:"type:varchar(100);unique;not null"`
	Status         RobotStatus    `gorm:"type:varchar(20);not null;default:'IDLE'"`
	BatteryLevel   int            `gorm:"type:integer;not null;default:100"`
	AvailableSlots int            `gorm:"type:integer;not null;default:4"`
	// Location is handled as PostGIS Geography
	Location       string         `gorm:"type:geography(POINT,4326)"` 
	HardwareStatus string         `gorm:"type:varchar(20);default:'OK'"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

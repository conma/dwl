package entity

import "github.com/google/uuid"

type Usr struct {
	Id *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (Usr) TableName() string {
	return "usr"
}

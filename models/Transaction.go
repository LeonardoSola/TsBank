package models

import (
	"time"
	"tsbank/infra/database"
)

type Transaction struct {
	ID            uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Value         int64     `json:"value" gorm:"not null"`
	OriginID      uint64    `json:"origin_id" gorm:"foreignKey:OriginID;not null"`
	Origin        *User     `json:"origin" gorm:"foreignKey:OriginID;<-:false"`
	DestinationID uint64    `json:"destination_id" gorm:"foreignKey:DestinationID;not null"`
	Destination   *User     `json:"destination" gorm:"foreignKey:DestinationID;<-:false"`
	Authorized    bool      `json:"authorized" gorm:"not null;default:false"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"not null;default:now()"`
}

func (transaction *Transaction) Create() error {
	return database.DB.Create(transaction).Error
}

func (transaction *Transaction) FindById() error {
	return database.DB.Where("id = ?", transaction.ID).First(transaction).Error
}

func (transaction *Transaction) Update() error {
	return database.DB.Model(transaction).Updates(transaction).Error
}

func (transaction *Transaction) Delete() error {
	return database.DB.Delete(transaction).Error
}
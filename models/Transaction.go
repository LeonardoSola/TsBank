package models

import (
	"time"
	"tsbank/infra/database"
)

type Transaction struct {
	ID            uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Value         int64     `json:"value" gorm:"not null"`
	OriginID      uint64    `json:"origin_id" gorm:"foreignKey:OriginID;not null"`
	Origin        *User     `json:"origin,omitempty" gorm:"foreignKey:OriginID;<-:false"`
	DestinationID uint64    `json:"destination_id" gorm:"foreignKey:DestinationID;not null"`
	Destination   *User     `json:"destination,omitempty" gorm:"foreignKey:DestinationID;<-:false"`
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

func (transaction *Transaction) FindAll(userID uint64, pagination *Pagination) (transactions []Transaction, err error) {
	err = database.DB.Where("origin_id = ? OR destination_id = ?", userID, userID).Limit(pagination.Limit).Offset(pagination.Offset).Find(&transactions).Error
	return transactions, err
}

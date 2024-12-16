package models

import (
	"github.com/google/uuid"
)

type CustomerModel struct {
	Id                  uuid.UUID
	UserId              uuid.UUID
	PreferredCategories []string
	//пока хз какие поля нужны
}

func (CustomerModel) TableName() string {
	return "customers"
}

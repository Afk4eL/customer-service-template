package dto

import "github.com/google/uuid"

type CustomerDto struct {
	Id                  uuid.UUID
	UserId              uuid.UUID
	PreferredCategories []string
	//пока хз какие поля нужны
}

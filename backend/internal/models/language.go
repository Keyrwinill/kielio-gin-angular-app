package models

import "github.com/google/uuid"

type Language struct {
	Oid  uuid.UUID `json:"oid"'`
	Name string    `json:"name"`
}

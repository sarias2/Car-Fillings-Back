package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Filling struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Km     string             `json:"name,omitempty" validate:"required"`
	Volume string             `json:"location,omitempty" validate:"required"`
	Date   string             `json:"title,omitempty" validate:"required"`
}

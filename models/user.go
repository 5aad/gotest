package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    UserName    string             `json:"userName,omitempty" bson:"userName,omitempty"`
}
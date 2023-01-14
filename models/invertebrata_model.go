package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Invertebrata struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nama             string             `json:"nama,omitempty" validate:"required"`
	Lokasi_Ditemukan string             `json:"lokasi_ditemukan,omitempty" validate:"required"`
	Waktu_Ditemukan  string             `json:"waktu_ditemukan,omitempty" validate:"required"`
}

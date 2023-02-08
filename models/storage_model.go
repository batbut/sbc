package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Storage struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Storage struct {
		RuangStorage string `bson:"Ruang_Storage" json:"ruang_storage" validate:"required"`
		Lantai       string `bson:"Lantai" json:"lantai" validate:"required"`
		Lajur        string `bson:"Lajur" json:"lajur" validate:"required"`
		Lemari       string `bson:"Lemari" json:"lemari" validate:"required"`
		Laci         string `bson:"Laci" json:"laci" validate:"required"`
		Slot         string `bson:"Slot" json:"slot" validate:"required"`
	} `bson:"Storage" json:"storage"`
	NonStorage struct {
		NamaNonStorage string `bson:"Nama_Non_Storage" json:"nama_non_storage" validate:"required"`
	} `bson:"Non_Storage" json:"non_storage"`
}

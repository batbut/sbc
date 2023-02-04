package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bmn struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NoRegister     string             `bson:"No Register" json:"no_register"`
	KategoriBMN    string             `bson:"Kategori BMN" json:"kategori_bmn"`
	TipeBMN        string             `bson:"Tipe BMN" json:"tipe_bmn"`
	NilaiPerolehan string             `bson:"Nilai Perolehan" json:"nilai_perolehan"`
	NilaiBuku      string             `bson:"Nilai Buku" json:"nilai_buku"`
}

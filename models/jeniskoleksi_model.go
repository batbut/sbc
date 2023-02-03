package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type JenisKoleksi struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaKoleksi       string             `bson:"Nama_Koleksi" json:"nama_koleksi"`
	JenisKoleksiFosil struct {
		JenisKoleksi     string `bson:"Jenis_Koleksi" json:"jenis_koleksi" validate:"required"`
		SubJenisKoleksi  string `bson:"Sub_Jenis_Koleksi" json:"sub_jenis_koleksi" validate:"required"`
		KodeJenisKoleksi string `bson:"Kode_Jenis_Koleksi" json:"kode_jenis_koleksi" validate:"required"`
	} `bson:"Jenis_Koleksi_Fosil" json:"jenis_koleksi_fosil"`
	JenisKoleksiBatuan struct {
		JenisKoleksi     string `bson:"Jenis_Koleksi" json:"jenis_koleksi" validate:"required"`
		SubJenisKoleksi  string `bson:"Sub_Jenis_Koleksi" json:"sub_jenis_koleksi" validate:"required"`
		KodeJenisKoleksi string `bson:"Kode_Jenis_Koleksi" json:"kode_jenis_koleksi" validate:"required"`
	} `bson:"Jenis_Koleksi_Batuan" json:"jenis_koleksi_batuan"`
}

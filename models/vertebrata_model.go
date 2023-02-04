package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vertebrata struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	No struct {
		Registrasi    string `bson:"Registrasi" json:"registrasi" validate:"required"`
		Inventarisasi string `bson:"Inventarisasi" json:"inventarisasi" validate:"required"`
		Laci          string `bson:"Laci" json:"laci" validate:"required"`
		Kotak         string `bson:"Kotak" json:"kotak" validate:"required"`
		KoleksiBaru   string `bson:"Koleksi Baru" json:"koleksi_baru" validate:"required"`
		KoleksiLama   string `bson:"Koleksi Lama" json:"koleksi_lama" validate:"required"`
	} `bson:"No" json:"no"`
	Pulau              string `bson:"Pulau" json:"pulau"`
	Spesies            string `bson:"Sepesies" json:"spesies"`
	Famili             string `bson:"Famili" json:"famili"`
	JenisKoleksi       string `bson:"Jenis Koleksi" json:"jenis_koleksi"`
	Determinasi        string `bson:"Determinasi" json:"determinasi"`
	Spesimen           string `bson:"Spesimen" json:"spesimen"`
	TipeKoleksi        string `bson:"Tipe Koleksi" json:"tipe_koleksi"`
	JumlahUtuh         string `bson:"Jumlah Utuh" json:"jumlah_utuh"`
	JumlahPecahan      string `bson:"Jumlah Pecahan" json:"jumlah_pecahan"`
	JumlahGabungan     string `bson:"Jumlah Gabungan" json:"jumlah_gabungan"`
	Lokasi             string `bson:"Lokasi" json:"lokasi"`
	KoordinatLokasi    string `bson:"Koordinat Lokasi" json:"koordinat_lokasi"`
	Formasi            string `bson:"Formasi" json:"formasi"`
	CaraPerolehan      string `bson:"Cara Perolehan" json:"cara_perolehan"`
	Umur               string `bson:"Umur" json:"umur"`
	ReferensiPublikasi string `bson:"Referensi Publikasi" json:"referensi_publikasi"`
	Kolektor           string `bson:"Kolektor" json:"kolektor"`
	TahunPenemuan      string `bson:"Tahun Penemuan" json:"tahun_penemuan"`
	Literatur          string `bson:"Literatur" json:"literatur"`
	Extra              string `bson:"Extra" json:"extra"`
	KondisiBenda       string `bson:"Kondisi Benda" json:"kondisi_benda"`
	Keterangan         string `bson:"Keterangan" json:"keterangan"`
	TanggalPencatatan  string `bson:"Tanggal Pencatatan" json:"tanggal_pencatatan"`
	Foto               string `bson:"Foto" json:"foto"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Batuan struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nomer struct {
		No_Reg  string `bson:"No_Reg" json:"No_Reg" validate:"required"`
		No_Inv  string `bson:"No_Inv" json:"No_Inv" validate:"required"`
		No_Awal string `bson:"No_Awal" json:"No_Awal" validate:"required"`
	} `bson:"Nomer" json:"Nomer"`
	Barang_Milik_Negara struct {
		Kode_Bmn string `bson:"Kode_Bmn" json:"Kode_Bmn" validate:"required"`
		Nup_Bmn  string `bson:"Nup_Bmn" json:"Nup_Bmn" validate:"required"`
		Merk_Bmn string `bson:"Merk_Bmn" json:"Merk_Bmn" validate:"required"`
	} `bson:"Barang_Milik_Negara" json:"Barang_Milik_Negara"`
	Determinator string `json:"determinator,omitempty" validate:"required"`
	Peta         struct {
		Nama_Peta    string `bson:"Nama_Peta" json:"Nama_Peta" validate:"required"`
		Skala_Peta   string `bson:"Skala_Peta" json:"Skala_Peta" validate:"required"`
		Koleksi_Peta string `bson:"Koleksi_Peta" json:"Koleksi_Peta" validate:"required"`
		Lembar_Peta  string `bson:"Lembar_Peta" json:"Lembar_Peta" validate:"required"`
	} `bson:"Peta" json:"Peta"`
	Cara_Perolehan string `bson:"Cara_Perolehan" json:"Cara_Perolehan" validate:"required"`
	Umur           string `json:"Umur" bson:"Umur" validate:"required"`
	Nama_Satuan    string `json:"Nama_Satuan" bson:"Nama_Satuan" validate:"required"`
	Kondisi        string `json:"Kondisi" bson:"Kondisi" validate:"required"`
	Dalam_Negeri   struct {
		Nama_Provinsi  string `bson:"Nama_Provinsi" json:"Nama_Provinsi" validate:"required"`
		Nama_Kabupaten string `bson:"Nama_Kabupaten" json:"Nama_Kabupaten" validate:"required"`
	} `bson:"Dalam_Negeri" json:"Dalam_Negeri"`
	Luar_Negeri struct {
		Keterangan_LN string `bson:"Keterangan_LN" json:"Keterangan_LN" validate:"required"`
	} `bson:"Luar_Negeri" json:"Luar_Negeri"`
	Koleksi struct {
		Nama_Koleksi       string `bson:"Nama_Koleksi" json:"Nama_Koleksi" validate:"required"`
		Jenis_Koleksi      string `bson:"Jenis_Koleksi" json:"Jenis_Koleksi" validate:"required"`
		Sub_Jenis_Koleksi  string `bson:"Sub_Jenis_Koleksi" json:"Sub_Jenis_Koleksi" validate:"required"`
		Kode_Jenis_Koleksi string `bson:"Kode_Jenis_Koleksi" json:"Kode_Jenis_Koleksi" validate:"required"`
		Deskripsi_Koleksi  string `bson:"Deskripsi_Koleksi" json:"Deskripsi_Koleksi" validate:"required"`
		Kelompok_Koleksi   string `bson:"Kelompok_Koleksi" json:"Kelompok_Koleksi" validate:"required"`
	} `bson:"Koleksi" json:"Koleksi"`
	Lokasi_Storage struct {
		Ruang_Storage string `bson:"Ruang_Storage" json:"Ruang_Storage" validate:"required"`
		Lantai        string `bson:"Lantai" json:"Lantai" validate:"required"`
		Lajur         string `bson:"Lajur" json:"Lajur" validate:"required"`
		Lemari        string `bson:"Lemari" json:"Lemari" validate:"required"`
		Laci          string `bson:"Laci" json:"Laci" validate:"required"`
		Slot          string `bson:"Slot" json:"Slot" validate:"required"`
	} `bson:"Lokasi_Storage" json:"Lokasi_Storage"`
	Lokasi_Non_Storage struct {
		Nama_Non_Storage string `bson:"Nama_Non_Storage" json:"Nama_Non_Storage" validate:"required"`
	} `bson:"Lokasi_Non_Storage" json:"Lokasi_Non_Storage"`
	Nama_Formasi     string `bson:"Nama_Formasi" json:"Nama_Formasi" validate:"required"`
	Keterangan       string `bson:"Keterangan" json:"Keterangan" validate:"required"`
	Pulau            string `bson:"Pulau" json:"Pulau" validate:"required"`
	Alamat_Lengkap   string `bson:"Alamat_Lengkap" json:"Alamat_Lengkap" validate:"required"`
	Koordinat_X      string `bson:"Koordinat_X" json:"Koordinat_X" validate:"required"`
	Koordinat_Y      string `bson:"Koordinat_Y" json:"Koordinat_Y" validate:"required"`
	Koordinat_Z      string `bson:"Koordinat_Z" json:"Koordinat_Z" validate:"required"`
	Tahun_Perolehan  string `bson:"Tahun_Perolehan" json:"Tahun_Perolehan" validate:"required"`
	Kolektor         string `bson:"Kolektor" json:"Kolektor" validate:"required"`
	Publikasi        string `bson:"Publikasi" json:"Publikasi" validate:"required"`
	Kepemilikan_Awal string `bson:"Kepemilikan_Awal" json:"Kepemilikan_Awal" validate:"required"`
	URL              string `bson:"URL" json:"URL" validate:"required"`
	Nilai_Perolehan  string `bson:"Nilai_Perolehan" json:"Nilai_Perolehan" validate:"required"`
	Nilai_Buku       string `bson:"Nilai_Buku" json:"Nilai_Buku" validate:"required"`
	Gambar_1         string `bson:"Gambar_1" json:"Gambar_1" validate:"required"`
	Gambar_2         string `bson:"Gambar_2" json:"Gambar_2" validate:"required"`
	Gambar_3         string `bson:"Gambar_3" json:"Gambar_3" validate:"required"`
}

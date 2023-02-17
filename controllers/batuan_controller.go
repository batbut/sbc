package controllers

import (
	"context"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var batuanCollection *mongo.Collection = configs.GetCollection(configs.DB, "batuan")
var validate_batuan = validator.New()

func CreateBatuan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var batuan models.Batuan
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&batuan); err != nil {
			c.JSON(http.StatusBadRequest, responses.BatuanResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_batuan.Struct(&batuan); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BatuanResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newBatuan := models.Batuan{
			Id: primitive.NewObjectID(),
			Nomer: struct {
				No_Reg  string "bson:\"No_Reg\" json:\"No_Reg\" validate:\"required\""
				No_Inv  string "bson:\"No_Inv\" json:\"No_Inv\" validate:\"required\""
				No_Awal string "bson:\"No_Awal\" json:\"No_Awal\" validate:\"required\""
			}{
				No_Reg:  batuan.Nomer.No_Reg,
				No_Inv:  batuan.Nomer.No_Inv,
				No_Awal: batuan.Nomer.No_Awal,
			},
			Badan_Milik_Negara: struct {
				Kode_Bmn string "bson:\"Kode_Bmn\" json:\"Kode_Bmn\" validate:\"required\""
				Nup_Bmn  string "bson:\"Nup_Bmn\" json:\"Nup_Bmn\" validate:\"required\""
				Merk_Bmn string "bson:\"Merk_Bmn\" json:\"Merk_Bmn\" validate:\"required\""
			}{
				Kode_Bmn: batuan.Badan_Milik_Negara.Kode_Bmn,
				Nup_Bmn:  batuan.Badan_Milik_Negara.Nup_Bmn,
				Merk_Bmn: batuan.Badan_Milik_Negara.Merk_Bmn,
			},
			Determinator: batuan.Determinator,
			Peta: struct {
				Nama_Peta    string "bson:\"Nama_Peta\" json:\"Nama_Peta\" validate:\"required\""
				Skala_Peta   string "bson:\"Skala_Peta\" json:\"Skala_Peta\" validate:\"required\""
				Koleksi_Peta string "bson:\"Koleksi_Peta\" json:\"Koleksi_Peta\" validate:\"required\""
				Lembar_Peta  string "bson:\"Lembar_Peta\" json:\"Lembar_Peta\" validate:\"required\""
			}{
				Nama_Peta:    batuan.Peta.Nama_Peta,
				Skala_Peta:   batuan.Peta.Skala_Peta,
				Koleksi_Peta: batuan.Peta.Koleksi_Peta,
				Lembar_Peta:  batuan.Peta.Lembar_Peta,
			},
			Cara_Perolehan: batuan.Cara_Perolehan,
			Umur:           batuan.Umur,
			Nama_Satuan:    batuan.Nama_Satuan,
			Kondisi:        batuan.Kondisi,
			Dalam_Negri: struct {
				Nama_Provinsi  string "bson:\"Nama_Provinsi\" json:\"Nama_Provinsi\" validate:\"required\""
				Nama_Kabupaten string "bson:\"Nama_Kabupaten\" json:\"Nama_Kabupaten\" validate:\"required\""
			}{
				Nama_Provinsi:  batuan.Dalam_Negri.Nama_Provinsi,
				Nama_Kabupaten: batuan.Dalam_Negri.Nama_Kabupaten,
			},
			Luar_Negri: struct {
				Keterangan_LN string "bson:\"Keterangan_LN\" json:\"Keterangan_LN\" validate:\"required\""
			}{
				Keterangan_LN: batuan.Luar_Negri.Keterangan_LN,
			},
			Koleksi: struct {
				Nama_Koleksi       string "bson:\"Nama_Koleksi\" json:\"Nama_Koleksi\" validate:\"required\""
				Jenis_Koleksi      string "bson:\"Jenis_Koleksi\" json:\"Jenis_Koleksi\" validate:\"required\""
				Sub_Jenis_Koleksi  string "bson:\"Sub_Jenis_Koleksi\" json:\"Sub_Jenis_Koleksi\" validate:\"required\""
				Kode_Jenis_Koleksi string "bson:\"Kode_Jenis_Koleksi\" json:\"Kode_Jenis_Koleksi\" validate:\"required\""
				Deskripsi_Koleksi  string "bson:\"Deskripsi_Koleksi\" json:\"Deskripsi_Koleksi\" validate:\"required\""
				Kelompok_Koleksi   string "bson:\"Kelompok_Koleksi\" json:\"Kelompok_Koleksi\" validate:\"required\""
			}{
				Nama_Koleksi:       batuan.Koleksi.Nama_Koleksi,
				Jenis_Koleksi:      batuan.Koleksi.Jenis_Koleksi,
				Sub_Jenis_Koleksi:  batuan.Koleksi.Sub_Jenis_Koleksi,
				Kode_Jenis_Koleksi: batuan.Koleksi.Kode_Jenis_Koleksi,
				Deskripsi_Koleksi:  batuan.Koleksi.Deskripsi_Koleksi,
				Kelompok_Koleksi:   batuan.Koleksi.Kelompok_Koleksi,
			},
			Lokasi_Storage: struct {
				Ruang_Storage string "bson:\"Ruang_Storage\" json:\"Ruang_Storage\" validate:\"required\""
				Lantai        string "bson:\"Lantai\" json:\"Lantai\" validate:\"required\""
				Lajur         string "bson:\"Lajur\" json:\"Lajur\" validate:\"required\""
				Lemari        string "bson:\"Lemari\" json:\"Lemari\" validate:\"required\""
				Laci          string "bson:\"Laci\" json:\"Laci\" validate:\"required\""
				Slot          string "bson:\"Slot\" json:\"Slot\" validate:\"required\""
			}{
				Ruang_Storage: batuan.Lokasi_Storage.Ruang_Storage,
				Lantai:        batuan.Lokasi_Storage.Lantai,
				Lajur:         batuan.Lokasi_Storage.Lajur,
				Lemari:        batuan.Lokasi_Storage.Lemari,
				Laci:          batuan.Lokasi_Storage.Laci,
				Slot:          batuan.Lokasi_Storage.Slot,
			},
			Lokasi_Non_Storage: struct {
				Nama_Non_Storage string "bson:\"Nama_Non_Storage\" json:\"Nama_Non_Storage\" validate:\"required\""
			}{
				Nama_Non_Storage: batuan.Lokasi_Non_Storage.Nama_Non_Storage,
			},
			Nama_Formasi:     batuan.Nama_Formasi,
			Keterangan:       batuan.Keterangan,
			Pulau:            batuan.Pulau,
			Alamat_Lengkap:   batuan.Alamat_Lengkap,
			Koordinat_X:      batuan.Koordinat_X,
			Koordinat_Y:      batuan.Koordinat_Y,
			Koordinat_Z:      batuan.Koordinat_Z,
			Tahun_Perolehan:  batuan.Tahun_Perolehan,
			Kolektor:         batuan.Kolektor,
			Publikasi:        batuan.Publikasi,
			Kepemilikan_Awal: batuan.Kepemilikan_Awal,
			URL:              batuan.URL,
			Nilai_Perolehan:  batuan.Nilai_Perolehan,
			Nilai_Buku:       batuan.Nilai_Buku,
			Gambar_1:         batuan.Gambar_1,
			Gambar_2:         batuan.Gambar_2,
			Gambar_3:         batuan.Gambar_3,
		}

		result, err := batuanCollection.InsertOne(ctx, newBatuan)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.BatuanResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetBatuan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		batuanId := c.Param("batuanId")
		var batuan models.Batuan
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(batuanId)

		err := batuanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&batuan)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.BatuanResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": batuan}})
	}
}

func EditBatuan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		batuanId := c.Param("batuanId")
		var batuan models.Batuan
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(batuanId)

		//validate the request body
		if err := c.BindJSON(&batuan); err != nil {
			c.JSON(http.StatusBadRequest, responses.BatuanResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_batuan.Struct(&batuan); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BatuanResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{
			"Nomer": bson.M{
				"No_Reg":  batuan.Nomer.No_Reg,
				"No_Inv":  batuan.Nomer.No_Inv,
				"No_Awal": batuan.Nomer.No_Awal,
			},
			"Badan_Milik_Negara": bson.M{
				"Kode_Bmn": batuan.Badan_Milik_Negara.Kode_Bmn,
				"Nup_Bmn":  batuan.Badan_Milik_Negara.Nup_Bmn,
				"Merk_Bmn": batuan.Badan_Milik_Negara.Merk_Bmn,
			},
			"Determinator": batuan.Determinator,
			"Peta": bson.M{
				"Nama_Peta":    batuan.Peta.Nama_Peta,
				"Skala_Peta":   batuan.Peta.Skala_Peta,
				"Koleksi_peta": batuan.Peta.Koleksi_Peta,
				"Lembar_Peta":  batuan.Peta.Lembar_Peta,
			},
			"Cara_Perolehan": batuan.Cara_Perolehan,
			"Umur":           batuan.Umur,
			"Nama_Satuan":    batuan.Nama_Satuan,
			"Kondisi":        batuan.Kondisi,
			"Dalam_Negri": bson.M{
				"Nama_Provinsi":  batuan.Dalam_Negri.Nama_Provinsi,
				"Nama_Kabupaten": batuan.Dalam_Negri.Nama_Kabupaten,
			},
			"Luar_Negri": bson.M{
				"Keterangan_LN": batuan.Luar_Negri.Keterangan_LN,
			},
			"Koleksi": bson.M{
				"Nama_Koleksi":       batuan.Koleksi.Nama_Koleksi,
				"Jenis_Koleksi":      batuan.Koleksi.Jenis_Koleksi,
				"Sub_Jenis_Koleksi":  batuan.Koleksi.Sub_Jenis_Koleksi,
				"Kode_Jenis_Koleksi": batuan.Koleksi.Kode_Jenis_Koleksi,
				"Kelompok_Koleksi":   batuan.Koleksi.Kelompok_Koleksi,
				"Deskripsi_Koleksi":  batuan.Koleksi.Deskripsi_Koleksi,
			},
			"Lokasi_Storage": bson.M{
				"Ruang_Storage": batuan.Lokasi_Storage.Ruang_Storage,
				"Lantai":        batuan.Lokasi_Storage.Lantai,
				"Lajur":         batuan.Lokasi_Storage.Lajur,
				"Lemari":        batuan.Lokasi_Storage.Lemari,
				"Laci":          batuan.Lokasi_Storage.Laci,
				"Slot":          batuan.Lokasi_Storage.Slot,
			},
			"Lokasi_Non_Storage": bson.M{
				"Nama_Non_Storage": batuan.Lokasi_Non_Storage.Nama_Non_Storage,
			},
			"Nama_Formasi":     batuan.Nama_Formasi,
			"Keterangan":       batuan.Keterangan,
			"Pulau":            batuan.Pulau,
			"Alamat_Lengkap":   batuan.Alamat_Lengkap,
			"Koordinat_X":      batuan.Koordinat_X,
			"Koordinat_Y":      batuan.Koordinat_Y,
			"Koordinat_Z":      batuan.Koordinat_Z,
			"Tahun_Perolehan":  batuan.Tahun_Perolehan,
			"Kolektor":         batuan.Kolektor,
			"Publikasi":        batuan.Publikasi,
			"Kepemilikan_Awal": batuan.Kepemilikan_Awal,
			"URL":              batuan.URL,
			"Nilai_Perolehan":  batuan.Nilai_Perolehan,
			"Nilai_Buku":       batuan.Nilai_Buku,
			"Gambar_1":         batuan.Gambar_1,
			"Gambar_2":         batuan.Gambar_2,
			"Gambar_3":         batuan.Gambar_3,
		}
		result, err := batuanCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated Batuan details
		var updatedBatuan models.Batuan
		if result.MatchedCount == 1 {
			err := batuanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedBatuan)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.BatuanResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedBatuan}})
	}
}

func DeleteBatuan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		batuanId := c.Param("batuanId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(batuanId)

		result, err := batuanCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.BatuanResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Batuan with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.BatuanResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Batuan successfully deleted!"}},
		)
	}
}

func GetAllBatuans() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var batuans []models.Batuan
		defer cancel()

		results, err := batuanCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleBatuan models.Batuan
			if err = results.Decode(&singleBatuan); err != nil {
				c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			batuans = append(batuans, singleBatuan)
		}

		c.JSON(http.StatusOK,
			responses.BatuanResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": batuans}},
		)
	}
}

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

var sumberdayageologiCollection *mongo.Collection = configs.GetCollection(configs.DB, "sumberdayageologi")
var validate_sumberdayageologi = validator.New()

func CreateSumberDayaGeologi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var sumberdayageologi models.SumberDayaGeologi
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&sumberdayageologi); err != nil {
			c.JSON(http.StatusBadRequest, responses.SumberDayaGeologiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_sumberdayageologi.Struct(&sumberdayageologi); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SumberDayaGeologiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newSumberDayaGeologi := models.SumberDayaGeologi{
			Id: primitive.NewObjectID(),
			Nomer: struct {
				No_Reg  string "bson:\"No_Reg\" json:\"No_Reg\" validate:\"required\""
				No_Inv  string "bson:\"No_Inv\" json:\"No_Inv\" validate:\"required\""
				No_Awal string "bson:\"No_Awal\" json:\"No_Awal\" validate:\"required\""
			}{
				No_Reg:  sumberdayageologi.Nomer.No_Reg,
				No_Inv:  sumberdayageologi.Nomer.No_Inv,
				No_Awal: sumberdayageologi.Nomer.No_Awal,
			},
			Badan_Milik_Negara: struct {
				Kode_Bmn string "bson:\"Kode_Bmn\" json:\"Kode_Bmn\" validate:\"required\""
				Nup_Bmn  string "bson:\"Nup_Bmn\" json:\"Nup_Bmn\" validate:\"required\""
				Merk_Bmn string "bson:\"Merk_Bmn\" json:\"Merk_Bmn\" validate:\"required\""
			}{
				Kode_Bmn: sumberdayageologi.Badan_Milik_Negara.Kode_Bmn,
				Nup_Bmn:  sumberdayageologi.Badan_Milik_Negara.Nup_Bmn,
				Merk_Bmn: sumberdayageologi.Badan_Milik_Negara.Merk_Bmn,
			},
			Determinator: sumberdayageologi.Determinator,
			Peta: struct {
				Nama_Peta    string "bson:\"Nama_Peta\" json:\"Nama_Peta\" validate:\"required\""
				Skala_Peta   string "bson:\"Skala_Peta\" json:\"Skala_Peta\" validate:\"required\""
				Koleksi_Peta string "bson:\"Koleksi_Peta\" json:\"Koleksi_Peta\" validate:\"required\""
				Lembar_Peta  string "bson:\"Lembar_Peta\" json:\"Lembar_Peta\" validate:\"required\""
			}{
				Nama_Peta:    sumberdayageologi.Peta.Nama_Peta,
				Skala_Peta:   sumberdayageologi.Peta.Skala_Peta,
				Koleksi_Peta: sumberdayageologi.Peta.Koleksi_Peta,
				Lembar_Peta:  sumberdayageologi.Peta.Lembar_Peta,
			},
			Cara_Perolehan: sumberdayageologi.Cara_Perolehan,
			Umur:           sumberdayageologi.Umur,
			Nama_Satuan:    sumberdayageologi.Nama_Satuan,
			Kondisi:        sumberdayageologi.Kondisi,
			Dalam_Negri: struct {
				Nama_Provinsi  string "bson:\"Nama_Provinsi\" json:\"Nama_Provinsi\" validate:\"required\""
				Nama_Kabupaten string "bson:\"Nama_Kabupaten\" json:\"Nama_Kabupaten\" validate:\"required\""
			}{
				Nama_Provinsi:  sumberdayageologi.Dalam_Negri.Nama_Provinsi,
				Nama_Kabupaten: sumberdayageologi.Dalam_Negri.Nama_Kabupaten,
			},
			Luar_Negri: struct {
				Keterangan_LN string "bson:\"Keterangan_LN\" json:\"Keterangan_LN\" validate:\"required\""
			}{
				Keterangan_LN: sumberdayageologi.Luar_Negri.Keterangan_LN,
			},
			Koleksi: struct {
				Nama_Koleksi       string "bson:\"Nama_Koleksi\" json:\"Nama_Koleksi\" validate:\"required\""
				Jenis_Koleksi      string "bson:\"Jenis_Koleksi\" json:\"Jenis_Koleksi\" validate:\"required\""
				Sub_Jenis_Koleksi  string "bson:\"Sub_Jenis_Koleksi\" json:\"Sub_Jenis_Koleksi\" validate:\"required\""
				Kode_Jenis_Koleksi string "bson:\"Kode_Jenis_Koleksi\" json:\"Kode_Jenis_Koleksi\" validate:\"required\""
				Deskripsi_Koleksi  string "bson:\"Deskripsi_Koleksi\" json:\"Deskripsi_Koleksi\" validate:\"required\""
				Kelompok_Koleksi   string "bson:\"Kelompok_Koleksi\" json:\"Kelompok_Koleksi\" validate:\"required\""
			}{
				Nama_Koleksi:       sumberdayageologi.Koleksi.Nama_Koleksi,
				Jenis_Koleksi:      sumberdayageologi.Koleksi.Jenis_Koleksi,
				Sub_Jenis_Koleksi:  sumberdayageologi.Koleksi.Sub_Jenis_Koleksi,
				Kode_Jenis_Koleksi: sumberdayageologi.Koleksi.Kode_Jenis_Koleksi,
				Deskripsi_Koleksi:  sumberdayageologi.Koleksi.Deskripsi_Koleksi,
				Kelompok_Koleksi:   sumberdayageologi.Koleksi.Kelompok_Koleksi,
			},
			Lokasi_Storage: struct {
				Ruang_Storage string "bson:\"Ruang_Storage\" json:\"Ruang_Storage\" validate:\"required\""
				Lantai        string "bson:\"Lantai\" json:\"Lantai\" validate:\"required\""
				Lajur         string "bson:\"Lajur\" json:\"Lajur\" validate:\"required\""
				Lemari        string "bson:\"Lemari\" json:\"Lemari\" validate:\"required\""
				Laci          string "bson:\"Laci\" json:\"Laci\" validate:\"required\""
				Slot          string "bson:\"Slot\" json:\"Slot\" validate:\"required\""
			}{
				Ruang_Storage: sumberdayageologi.Lokasi_Storage.Ruang_Storage,
				Lantai:        sumberdayageologi.Lokasi_Storage.Lantai,
				Lajur:         sumberdayageologi.Lokasi_Storage.Lajur,
				Lemari:        sumberdayageologi.Lokasi_Storage.Lemari,
				Laci:          sumberdayageologi.Lokasi_Storage.Laci,
				Slot:          sumberdayageologi.Lokasi_Storage.Slot,
			},
			Lokasi_Non_Storage: struct {
				Nama_Non_Storage string "bson:\"Nama_Non_Storage\" json:\"Nama_Non_Storage\" validate:\"required\""
			}{
				Nama_Non_Storage: sumberdayageologi.Lokasi_Non_Storage.Nama_Non_Storage,
			},
			Nama_Formasi:     sumberdayageologi.Nama_Formasi,
			Keterangan:       sumberdayageologi.Keterangan,
			Pulau:            sumberdayageologi.Pulau,
			Alamat_Lengkap:   sumberdayageologi.Alamat_Lengkap,
			Koordinat_X:      sumberdayageologi.Koordinat_X,
			Koordinat_Y:      sumberdayageologi.Koordinat_Y,
			Koordinat_Z:      sumberdayageologi.Koordinat_Z,
			Tahun_Perolehan:  sumberdayageologi.Tahun_Perolehan,
			Kolektor:         sumberdayageologi.Kolektor,
			Publikasi:        sumberdayageologi.Publikasi,
			Kepemilikan_Awal: sumberdayageologi.Kepemilikan_Awal,
			URL:              sumberdayageologi.URL,
			Nilai_Perolehan:  sumberdayageologi.Nilai_Perolehan,
			Nilai_Buku:       sumberdayageologi.Nilai_Buku,
			Gambar_1:         sumberdayageologi.Gambar_1,
			Gambar_2:         sumberdayageologi.Gambar_2,
			Gambar_3:         sumberdayageologi.Gambar_3,
		}

		result, err := sumberdayageologiCollection.InsertOne(ctx, newSumberDayaGeologi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.SumberDayaGeologiResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetSumberDayaGeologi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		sumberdayageologiId := c.Param("sumberdayageologiId")
		var sumberdayageologi models.SumberDayaGeologi
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(sumberdayageologiId)

		err := sumberdayageologiCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&sumberdayageologi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.SumberDayaGeologiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": sumberdayageologi}})
	}
}

func EditSumberDayaGeologi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		sumberdayageologiId := c.Param("sumberdayageologiId")
		var sumberdayageologi models.SumberDayaGeologi
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(sumberdayageologiId)

		//validate the request body
		if err := c.BindJSON(&sumberdayageologi); err != nil {
			c.JSON(http.StatusBadRequest, responses.SumberDayaGeologiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_sumberdayageologi.Struct(&sumberdayageologi); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SumberDayaGeologiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{
			"Nomer": bson.M{
				"No_Reg":  sumberdayageologi.Nomer.No_Reg,
				"No_Inv":  sumberdayageologi.Nomer.No_Inv,
				"No_Awal": sumberdayageologi.Nomer.No_Awal,
			},
			"Badan_Milik_Negara": bson.M{
				"Kode_Bmn": sumberdayageologi.Badan_Milik_Negara.Kode_Bmn,
				"Nup_Bmn":  sumberdayageologi.Badan_Milik_Negara.Nup_Bmn,
				"Merk_Bmn": sumberdayageologi.Badan_Milik_Negara.Merk_Bmn,
			},
			"Determinator": sumberdayageologi.Determinator,
			"Peta": bson.M{
				"Nama_Peta":    sumberdayageologi.Peta.Nama_Peta,
				"Skala_Peta":   sumberdayageologi.Peta.Skala_Peta,
				"Koleksi_peta": sumberdayageologi.Peta.Koleksi_Peta,
				"Lembar_Peta":  sumberdayageologi.Peta.Lembar_Peta,
			},
			"Cara_Perolehan": sumberdayageologi.Cara_Perolehan,
			"Umur":           sumberdayageologi.Umur,
			"Nama_Satuan":    sumberdayageologi.Nama_Satuan,
			"Kondisi":        sumberdayageologi.Kondisi,
			"Dalam_Negri": bson.M{
				"Nama_Provinsi":  sumberdayageologi.Dalam_Negri.Nama_Provinsi,
				"Nama_Kabupaten": sumberdayageologi.Dalam_Negri.Nama_Kabupaten,
			},
			"Luar_Negri": bson.M{
				"Keterangan_LN": sumberdayageologi.Luar_Negri.Keterangan_LN,
			},
			"Koleksi": bson.M{
				"Nama_Koleksi":       sumberdayageologi.Koleksi.Nama_Koleksi,
				"Jenis_Koleksi":      sumberdayageologi.Koleksi.Jenis_Koleksi,
				"Sub_Jenis_Koleksi":  sumberdayageologi.Koleksi.Sub_Jenis_Koleksi,
				"Kode_Jenis_Koleksi": sumberdayageologi.Koleksi.Kode_Jenis_Koleksi,
				"Kelompok_Koleksi":   sumberdayageologi.Koleksi.Kelompok_Koleksi,
				"Deskripsi_Koleksi":  sumberdayageologi.Koleksi.Deskripsi_Koleksi,
			},
			"Lokasi_Storage": bson.M{
				"Ruang_Storage": sumberdayageologi.Lokasi_Storage.Ruang_Storage,
				"Lantai":        sumberdayageologi.Lokasi_Storage.Lantai,
				"Lajur":         sumberdayageologi.Lokasi_Storage.Lajur,
				"Lemari":        sumberdayageologi.Lokasi_Storage.Lemari,
				"Laci":          sumberdayageologi.Lokasi_Storage.Laci,
				"Slot":          sumberdayageologi.Lokasi_Storage.Slot,
			},
			"Lokasi_Non_Storage": bson.M{
				"Nama_Non_Storage": sumberdayageologi.Lokasi_Non_Storage.Nama_Non_Storage,
			},
			"Nama_Formasi":     sumberdayageologi.Nama_Formasi,
			"Keterangan":       sumberdayageologi.Keterangan,
			"Pulau":            sumberdayageologi.Pulau,
			"Alamat_Lengkap":   sumberdayageologi.Alamat_Lengkap,
			"Koordinat_X":      sumberdayageologi.Koordinat_X,
			"Koordinat_Y":      sumberdayageologi.Koordinat_Y,
			"Koordinat_Z":      sumberdayageologi.Koordinat_Z,
			"Tahun_Perolehan":  sumberdayageologi.Tahun_Perolehan,
			"Kolektor":         sumberdayageologi.Kolektor,
			"Publikasi":        sumberdayageologi.Publikasi,
			"Kepemilikan_Awal": sumberdayageologi.Kepemilikan_Awal,
			"URL":              sumberdayageologi.URL,
			"Nilai_Perolehan":  sumberdayageologi.Nilai_Perolehan,
			"Nilai_Buku":       sumberdayageologi.Nilai_Buku,
			"Gambar_1":         sumberdayageologi.Gambar_1,
			"Gambar_2":         sumberdayageologi.Gambar_2,
			"Gambar_3":         sumberdayageologi.Gambar_3,
		}
		result, err := sumberdayageologiCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated SumberDayaGeologi details
		var updatedSumberDayaGeologi models.SumberDayaGeologi
		if result.MatchedCount == 1 {
			err := sumberdayageologiCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedSumberDayaGeologi)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.SumberDayaGeologiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedSumberDayaGeologi}})
	}
}

func DeleteSumberDayaGeologi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		sumberdayageologiId := c.Param("sumberdayageologiId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(sumberdayageologiId)

		result, err := sumberdayageologiCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.SumberDayaGeologiResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "SumberDayaGeologi with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.SumberDayaGeologiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "SumberDayaGeologi successfully deleted!"}},
		)
	}
}

func GetAllSumberDayaGeologis() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var sumberdayageologis []models.SumberDayaGeologi
		defer cancel()

		results, err := sumberdayageologiCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleSumberDayaGeologi models.SumberDayaGeologi
			if err = results.Decode(&singleSumberDayaGeologi); err != nil {
				c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			sumberdayageologis = append(sumberdayageologis, singleSumberDayaGeologi)
		}

		c.JSON(http.StatusOK,
			responses.SumberDayaGeologiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": sumberdayageologis}},
		)
	}
}

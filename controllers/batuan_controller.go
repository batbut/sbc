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
	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var batuanCollection *mongo.Collection = configs.GetCollection(configs.DB, "batuan")
var validate_batuan = validator.New()

// CreateBatuan godoc
// @Summary Create a new Batuan
// @Description Create a new Batuan with the input payload
// @Tags Batuan
// @Accept  json
// @Produce  json
// @Param batuan body models.Batuan true "The batuan to create"
// @Success 200 {object} responses.BatuanResponse
// @Failure 400 {object} responses.BatuanResponse
// @Router /batuan [post]
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
			Barang_Milik_Negara: struct {
				Kode_Bmn string "bson:\"Kode_Bmn\" json:\"Kode_Bmn\" validate:\"required\""
				Nup_Bmn  string "bson:\"Nup_Bmn\" json:\"Nup_Bmn\" validate:\"required\""
				Merk_Bmn string "bson:\"Merk_Bmn\" json:\"Merk_Bmn\" validate:\"required\""
			}{
				Kode_Bmn: batuan.Barang_Milik_Negara.Kode_Bmn,
				Nup_Bmn:  batuan.Barang_Milik_Negara.Nup_Bmn,
				Merk_Bmn: batuan.Barang_Milik_Negara.Merk_Bmn,
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
			Dalam_Negeri: struct {
				Nama_Provinsi  string "bson:\"Nama_Provinsi\" json:\"Nama_Provinsi\" validate:\"required\""
				Nama_Kabupaten string "bson:\"Nama_Kabupaten\" json:\"Nama_Kabupaten\" validate:\"required\""
			}{
				Nama_Provinsi:  batuan.Dalam_Negeri.Nama_Provinsi,
				Nama_Kabupaten: batuan.Dalam_Negeri.Nama_Kabupaten,
			},
			Luar_Negeri: struct {
				Keterangan_LN string "bson:\"Keterangan_LN\" json:\"Keterangan_LN\" validate:\"required\""
			}{
				Keterangan_LN: batuan.Luar_Negeri.Keterangan_LN,
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

// GetBatuan godoc
// @Summary Get Batuan by ID
// @Description Get a Batuan by its ID
// @Tags Batuan
// @ID get-batuan-by-id
// @Produce json
// @Param batuanId path string true "Batuan ID"
// @Success 200 {object} responses.BatuanResponse
// @Failure 500 {object} responses.BatuanResponse
// @Router /batuan/{batuanId} [get]
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

// EditBatuan edits an existing Batuan.
// @Summary Edit an existing Batuan
// @Description Edit an existing Batuan
// @Tags Batuan
// @Accept json
// @Produce json
// @Param batuanId path string true "ID of the Batuan to edit"
// @Param body body models.Batuan true "Batuan object that needs to be edited"
// @Success 200 {object} responses.BatuanResponse
// @Failure 400 {object} responses.BatuanResponse
// @Router /batuan/{batuanId} [put]
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
				"Kode_Bmn": batuan.Barang_Milik_Negara.Kode_Bmn,
				"Nup_Bmn":  batuan.Barang_Milik_Negara.Nup_Bmn,
				"Merk_Bmn": batuan.Barang_Milik_Negara.Merk_Bmn,
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
			"Dalam_Negeri": bson.M{
				"Nama_Provinsi":  batuan.Dalam_Negeri.Nama_Provinsi,
				"Nama_Kabupaten": batuan.Dalam_Negeri.Nama_Kabupaten,
			},
			"Luar_Negeri": bson.M{
				"Keterangan_LN": batuan.Luar_Negeri.Keterangan_LN,
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

// DeleteBatuan godoc
// @Summary Delete a batuan by ID
// @Description Delete a batuan by ID
// @Tags Batuan
// @Param batuanId path string true "ID of the batuan to delete"
// @Produce json
// @Success 200 {object} responses.BatuanResponse
// @Failure 404 {object} responses.BatuanResponse
// @Failure 500 {object} responses.BatuanResponse
// @Router /batuan/{batuanId} [delete]
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

// GetAllBatuans godoc
// @Summary Get all batuans
// @Description Retrieve all batuans from the database
// @Tags Batuan
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.BatuanResponse
// @Failure 500 {object} responses.BatuanResponse
// @Router /batuans [get]
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

// ExportBatuanToExcel export data batuan to excel.
// @Summary Export data batuan to excel
// @Description Get data batuan from MongoDB and export to excel
// @Tags Batuan
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Data batuan exported to excel successfully"
// @Failure 500 {object} responses.BatuanResponse
// @Router /batuans/export [get]
func ExportBatuanToExcel() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		file := xlsx.NewFile()
		sheet, err := file.AddSheet("Data Batuan")
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		rows, err := batuanCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// add headers
		row := sheet.AddRow()
		row.AddCell().Value = "No Register"
		row.AddCell().Value = "No Inventaris"
		row.AddCell().Value = "No Awal"
		row.AddCell().Value = "Kode BMN"
		row.AddCell().Value = "Nup BMN"
		row.AddCell().Value = "Merk BMN"
		row.AddCell().Value = "Determinator"
		row.AddCell().Value = "Nama Peta"
		row.AddCell().Value = "Skala Peta"
		row.AddCell().Value = "Koleksi Peta"
		row.AddCell().Value = "Lembar Peta"
		row.AddCell().Value = "Cara Perolehan"
		row.AddCell().Value = "Umur"
		row.AddCell().Value = "Nama Satuan"
		row.AddCell().Value = "Kondisi"
		row.AddCell().Value = "Nama Provinsi"
		row.AddCell().Value = "Nama Kabupaten"
		row.AddCell().Value = "Keterangan Luar Negeri"
		row.AddCell().Value = "Nama Koleksi"
		row.AddCell().Value = "Jenis Koleksi"
		row.AddCell().Value = "Sub Jenis Koleksi"
		row.AddCell().Value = "Kode Jenis Koleksi"
		row.AddCell().Value = "Deskripsi Koleksi"
		row.AddCell().Value = "Kelompok Koleksi"
		row.AddCell().Value = "Ruang Storage"
		row.AddCell().Value = "Lantai"
		row.AddCell().Value = "Lajur"
		row.AddCell().Value = "Lemari"
		row.AddCell().Value = "Laci"
		row.AddCell().Value = "Slot"
		row.AddCell().Value = "Nama Non Storage"
		row.AddCell().Value = "Nama Formasi"
		row.AddCell().Value = "Keterangan"
		row.AddCell().Value = "Pulau"
		row.AddCell().Value = "Alamat Lengkap"
		row.AddCell().Value = "Koordinat X"
		row.AddCell().Value = "Koordinat Y"
		row.AddCell().Value = "Koordinat Z"
		row.AddCell().Value = "Tahun Perolehan"
		row.AddCell().Value = "Kolektor"
		row.AddCell().Value = "Publikasi"
		row.AddCell().Value = "Kepemilikan Awal"
		row.AddCell().Value = "URL"
		row.AddCell().Value = "Nilai Perolehan"
		row.AddCell().Value = "Nilai Buku"
		row.AddCell().Value = "Gambar 1"
		row.AddCell().Value = "Gambar 2"
		row.AddCell().Value = "Gambar 3"

		for rows.Next(ctx) {
			var batuan models.Batuan
			if err := rows.Decode(&batuan); err != nil {
				c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			row := sheet.AddRow()
			row.AddCell().Value = batuan.Nomer.No_Reg
			row.AddCell().Value = batuan.Nomer.No_Inv
			row.AddCell().Value = batuan.Nomer.No_Awal
			row.AddCell().Value = batuan.Barang_Milik_Negara.Kode_Bmn
			row.AddCell().Value = batuan.Barang_Milik_Negara.Nup_Bmn
			row.AddCell().Value = batuan.Barang_Milik_Negara.Merk_Bmn
			row.AddCell().Value = batuan.Determinator
			row.AddCell().Value = batuan.Peta.Nama_Peta
			row.AddCell().Value = batuan.Peta.Skala_Peta
			row.AddCell().Value = batuan.Peta.Koleksi_Peta
			row.AddCell().Value = batuan.Peta.Lembar_Peta
			row.AddCell().Value = batuan.Cara_Perolehan
			row.AddCell().Value = batuan.Umur
			row.AddCell().Value = batuan.Nama_Satuan
			row.AddCell().Value = batuan.Kondisi
			row.AddCell().Value = batuan.Dalam_Negeri.Nama_Provinsi
			row.AddCell().Value = batuan.Dalam_Negeri.Nama_Kabupaten
			row.AddCell().Value = batuan.Luar_Negeri.Keterangan_LN
			row.AddCell().Value = batuan.Koleksi.Nama_Koleksi
			row.AddCell().Value = batuan.Koleksi.Jenis_Koleksi
			row.AddCell().Value = batuan.Koleksi.Sub_Jenis_Koleksi
			row.AddCell().Value = batuan.Koleksi.Kode_Jenis_Koleksi
			row.AddCell().Value = batuan.Koleksi.Deskripsi_Koleksi
			row.AddCell().Value = batuan.Koleksi.Kelompok_Koleksi
			row.AddCell().Value = batuan.Lokasi_Storage.Ruang_Storage
			row.AddCell().Value = batuan.Lokasi_Storage.Lantai
			row.AddCell().Value = batuan.Lokasi_Storage.Lajur
			row.AddCell().Value = batuan.Lokasi_Storage.Lemari
			row.AddCell().Value = batuan.Lokasi_Storage.Laci
			row.AddCell().Value = batuan.Lokasi_Storage.Slot
			row.AddCell().Value = batuan.Lokasi_Non_Storage.Nama_Non_Storage
			row.AddCell().Value = batuan.Nama_Formasi
			row.AddCell().Value = batuan.Keterangan
			row.AddCell().Value = batuan.Pulau
			row.AddCell().Value = batuan.Alamat_Lengkap
			row.AddCell().Value = batuan.Koordinat_X
			row.AddCell().Value = batuan.Koordinat_Y
			row.AddCell().Value = batuan.Koordinat_Z
			row.AddCell().Value = batuan.Tahun_Perolehan
			row.AddCell().Value = batuan.Kolektor
			row.AddCell().Value = batuan.Publikasi
			row.AddCell().Value = batuan.Kepemilikan_Awal
			row.AddCell().Value = batuan.URL
			row.AddCell().Value = batuan.Nilai_Perolehan
			row.AddCell().Value = batuan.Nilai_Buku
			row.AddCell().Value = batuan.Gambar_1
			row.AddCell().Value = batuan.Gambar_2
			row.AddCell().Value = batuan.Gambar_3
		}
		filename := "data_batuan.xlsx"
		err = file.Save(filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BatuanResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")
		c.File(filename)
	}
}

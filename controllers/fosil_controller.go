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

var fosilCollection *mongo.Collection = configs.GetCollection(configs.DB, "fosil")
var validate_fosil = validator.New()

// CreateFosil godoc
// @Summary Create a new Fosil
// @Description Create a new Fosil with the input payload
// @Tags Fosil
// @Accept  json
// @Produce  json
// @Param fosil body models.Fosil true "The fosil to create"
// @Success 200 {object} responses.FosilResponse
// @Failure 400 {object} responses.FosilResponse
// @Router /fosil [post]
func CreateFosil() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var fosil models.Fosil
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&fosil); err != nil {
			c.JSON(http.StatusBadRequest, responses.FosilResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_fosil.Struct(&fosil); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.FosilResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newFosil := models.Fosil{
			Id: primitive.NewObjectID(),
			Nomer: struct {
				No_Reg  string "bson:\"No_Reg\" json:\"No_Reg\" validate:\"required\""
				No_Inv  string "bson:\"No_Inv\" json:\"No_Inv\" validate:\"required\""
				No_Awal string "bson:\"No_Awal\" json:\"No_Awal\" validate:\"required\""
			}{
				No_Reg:  fosil.Nomer.No_Reg,
				No_Inv:  fosil.Nomer.No_Inv,
				No_Awal: fosil.Nomer.No_Awal,
			},
			Barang_Milik_Negara: struct {
				Kode_Bmn string "bson:\"Kode_Bmn\" json:\"Kode_Bmn\" validate:\"required\""
				Nup_Bmn  string "bson:\"Nup_Bmn\" json:\"Nup_Bmn\" validate:\"required\""
				Merk_Bmn string "bson:\"Merk_Bmn\" json:\"Merk_Bmn\" validate:\"required\""
			}{
				Kode_Bmn: fosil.Barang_Milik_Negara.Kode_Bmn,
				Nup_Bmn:  fosil.Barang_Milik_Negara.Nup_Bmn,
				Merk_Bmn: fosil.Barang_Milik_Negara.Merk_Bmn,
			},
			Determinator: fosil.Determinator,
			Peta: struct {
				Nama_Peta    string "bson:\"Nama_Peta\" json:\"Nama_Peta\" validate:\"required\""
				Skala_Peta   string "bson:\"Skala_Peta\" json:\"Skala_Peta\" validate:\"required\""
				Koleksi_Peta string "bson:\"Koleksi_Peta\" json:\"Koleksi_Peta\" validate:\"required\""
				Lembar_Peta  string "bson:\"Lembar_Peta\" json:\"Lembar_Peta\" validate:\"required\""
			}{
				Nama_Peta:    fosil.Peta.Nama_Peta,
				Skala_Peta:   fosil.Peta.Skala_Peta,
				Koleksi_Peta: fosil.Peta.Koleksi_Peta,
				Lembar_Peta:  fosil.Peta.Lembar_Peta,
			},
			Cara_Perolehan: fosil.Cara_Perolehan,
			Umur:           fosil.Umur,
			Nama_Satuan:    fosil.Nama_Satuan,
			Kondisi:        fosil.Kondisi,
			Dalam_Negeri: struct {
				Nama_Provinsi  string "bson:\"Nama_Provinsi\" json:\"Nama_Provinsi\" validate:\"required\""
				Nama_Kabupaten string "bson:\"Nama_Kabupaten\" json:\"Nama_Kabupaten\" validate:\"required\""
			}{
				Nama_Provinsi:  fosil.Dalam_Negeri.Nama_Provinsi,
				Nama_Kabupaten: fosil.Dalam_Negeri.Nama_Kabupaten,
			},
			Luar_Negeri: struct {
				Keterangan_LN string "bson:\"Keterangan_LN\" json:\"Keterangan_LN\" validate:\"required\""
			}{
				Keterangan_LN: fosil.Luar_Negeri.Keterangan_LN,
			},
			Koleksi: struct {
				Nama_Koleksi       string "bson:\"Nama_Koleksi\" json:\"Nama_Koleksi\" validate:\"required\""
				Jenis_Koleksi      string "bson:\"Jenis_Koleksi\" json:\"Jenis_Koleksi\" validate:\"required\""
				Sub_Jenis_Koleksi  string "bson:\"Sub_Jenis_Koleksi\" json:\"Sub_Jenis_Koleksi\" validate:\"required\""
				Kode_Jenis_Koleksi string "bson:\"Kode_Jenis_Koleksi\" json:\"Kode_Jenis_Koleksi\" validate:\"required\""
				Deskripsi_Koleksi  string "bson:\"Deskripsi_Koleksi\" json:\"Deskripsi_Koleksi\" validate:\"required\""
				Kelompok_Koleksi   string "bson:\"Kelompok_Koleksi\" json:\"Kelompok_Koleksi\" validate:\"required\""
			}{
				Nama_Koleksi:       fosil.Koleksi.Nama_Koleksi,
				Jenis_Koleksi:      fosil.Koleksi.Jenis_Koleksi,
				Sub_Jenis_Koleksi:  fosil.Koleksi.Sub_Jenis_Koleksi,
				Kode_Jenis_Koleksi: fosil.Koleksi.Kode_Jenis_Koleksi,
				Deskripsi_Koleksi:  fosil.Koleksi.Deskripsi_Koleksi,
				Kelompok_Koleksi:   fosil.Koleksi.Kelompok_Koleksi,
			},
			Lokasi_Storage: struct {
				Ruang_Storage string "bson:\"Ruang_Storage\" json:\"Ruang_Storage\" validate:\"required\""
				Lantai        string "bson:\"Lantai\" json:\"Lantai\" validate:\"required\""
				Lajur         string "bson:\"Lajur\" json:\"Lajur\" validate:\"required\""
				Lemari        string "bson:\"Lemari\" json:\"Lemari\" validate:\"required\""
				Laci          string "bson:\"Laci\" json:\"Laci\" validate:\"required\""
				Slot          string "bson:\"Slot\" json:\"Slot\" validate:\"required\""
			}{
				Ruang_Storage: fosil.Lokasi_Storage.Ruang_Storage,
				Lantai:        fosil.Lokasi_Storage.Lantai,
				Lajur:         fosil.Lokasi_Storage.Lajur,
				Lemari:        fosil.Lokasi_Storage.Lemari,
				Laci:          fosil.Lokasi_Storage.Laci,
				Slot:          fosil.Lokasi_Storage.Slot,
			},
			Lokasi_Non_Storage: struct {
				Nama_Non_Storage string "bson:\"Nama_Non_Storage\" json:\"Nama_Non_Storage\" validate:\"required\""
			}{
				Nama_Non_Storage: fosil.Lokasi_Non_Storage.Nama_Non_Storage,
			},
			Nama_Formasi:     fosil.Nama_Formasi,
			Keterangan:       fosil.Keterangan,
			Pulau:            fosil.Pulau,
			Alamat_Lengkap:   fosil.Alamat_Lengkap,
			Koordinat_X:      fosil.Koordinat_X,
			Koordinat_Y:      fosil.Koordinat_Y,
			Koordinat_Z:      fosil.Koordinat_Z,
			Tahun_Perolehan:  fosil.Tahun_Perolehan,
			Kolektor:         fosil.Kolektor,
			Publikasi:        fosil.Publikasi,
			Kepemilikan_Awal: fosil.Kepemilikan_Awal,
			URL:              fosil.URL,
			Nilai_Perolehan:  fosil.Nilai_Perolehan,
			Nilai_Buku:       fosil.Nilai_Buku,
			Gambar_1:         fosil.Gambar_1,
			Gambar_2:         fosil.Gambar_2,
			Gambar_3:         fosil.Gambar_3,
		}

		result, err := fosilCollection.InsertOne(ctx, newFosil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.FosilResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// GetFosil godoc
// @Summary Get Fosil by ID
// @Description Get a Fosil by its ID
// @Tags Fosil
// @ID get-fosil-by-id
// @Produce json
// @Param fosilId path string true "Fosil ID"
// @Success 200 {object} responses.FosilResponse
// @Failure 500 {object} responses.FosilResponse
// @Router /fosil/{fosilId} [get]
func GetFosil() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		fosilId := c.Param("fosilId")
		var fosil models.Fosil
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(fosilId)

		err := fosilCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&fosil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.FosilResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": fosil}})
	}
}

// EditFosil edits an existing Fosil.
// @Summary Edit an existing Fosil
// @Description Edit an existing Fosil
// @Tags Fosil
// @Accept json
// @Produce json
// @Param fosilId path string true "ID of the Fosil to edit"
// @Param body body models.Fosil true "Fosil object that needs to be edited"
// @Success 200 {object} responses.FosilResponse
// @Failure 400 {object} responses.FosilResponse
// @Router /fosil/{fosilId} [put]
func EditFosil() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		fosilId := c.Param("fosilId")
		var fosil models.Fosil
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(fosilId)

		//validate the request body
		if err := c.BindJSON(&fosil); err != nil {
			c.JSON(http.StatusBadRequest, responses.FosilResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_fosil.Struct(&fosil); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.FosilResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{
			"Nomer": bson.M{
				"No_Reg":  fosil.Nomer.No_Reg,
				"No_Inv":  fosil.Nomer.No_Inv,
				"No_Awal": fosil.Nomer.No_Awal,
			},
			"Barang_Milik_Negara": bson.M{
				"Kode_Bmn": fosil.Barang_Milik_Negara.Kode_Bmn,
				"Nup_Bmn":  fosil.Barang_Milik_Negara.Nup_Bmn,
				"Merk_Bmn": fosil.Barang_Milik_Negara.Merk_Bmn,
			},
			"Determinator": fosil.Determinator,
			"Peta": bson.M{
				"Nama_Peta":    fosil.Peta.Nama_Peta,
				"Skala_Peta":   fosil.Peta.Skala_Peta,
				"Koleksi_peta": fosil.Peta.Koleksi_Peta,
				"Lembar_Peta":  fosil.Peta.Lembar_Peta,
			},
			"Cara_Perolehan": fosil.Cara_Perolehan,
			"Umur":           fosil.Umur,
			"Nama_Satuan":    fosil.Nama_Satuan,
			"Kondisi":        fosil.Kondisi,
			"Dalam_Negeri": bson.M{
				"Nama_Provinsi":  fosil.Dalam_Negeri.Nama_Provinsi,
				"Nama_Kabupaten": fosil.Dalam_Negeri.Nama_Kabupaten,
			},
			"Luar_Negeri": bson.M{
				"Keterangan_LN": fosil.Luar_Negeri.Keterangan_LN,
			},
			"Koleksi": bson.M{
				"Nama_Koleksi":       fosil.Koleksi.Nama_Koleksi,
				"Jenis_Koleksi":      fosil.Koleksi.Jenis_Koleksi,
				"Sub_Jenis_Koleksi":  fosil.Koleksi.Sub_Jenis_Koleksi,
				"Kode_Jenis_Koleksi": fosil.Koleksi.Kode_Jenis_Koleksi,
				"Kelompok_Koleksi":   fosil.Koleksi.Kelompok_Koleksi,
				"Deskripsi_Koleksi":  fosil.Koleksi.Deskripsi_Koleksi,
			},
			"Lokasi_Storage": bson.M{
				"Ruang_Storage": fosil.Lokasi_Storage.Ruang_Storage,
				"Lantai":        fosil.Lokasi_Storage.Lantai,
				"Lajur":         fosil.Lokasi_Storage.Lajur,
				"Lemari":        fosil.Lokasi_Storage.Lemari,
				"Laci":          fosil.Lokasi_Storage.Laci,
				"Slot":          fosil.Lokasi_Storage.Slot,
			},
			"Lokasi_Non_Storage": bson.M{
				"Nama_Non_Storage": fosil.Lokasi_Non_Storage.Nama_Non_Storage,
			},
			"Nama_Formasi":     fosil.Nama_Formasi,
			"Keterangan":       fosil.Keterangan,
			"Pulau":            fosil.Pulau,
			"Alamat_Lengkap":   fosil.Alamat_Lengkap,
			"Koordinat_X":      fosil.Koordinat_X,
			"Koordinat_Y":      fosil.Koordinat_Y,
			"Koordinat_Z":      fosil.Koordinat_Z,
			"Tahun_Perolehan":  fosil.Tahun_Perolehan,
			"Kolektor":         fosil.Kolektor,
			"Publikasi":        fosil.Publikasi,
			"Kepemilikan_Awal": fosil.Kepemilikan_Awal,
			"URL":              fosil.URL,
			"Nilai_Perolehan":  fosil.Nilai_Perolehan,
			"Nilai_Buku":       fosil.Nilai_Buku,
			"Gambar_1":         fosil.Gambar_1,
			"Gambar_2":         fosil.Gambar_2,
			"Gambar_3":         fosil.Gambar_3,
		}
		result, err := fosilCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated Fosil details
		var updatedFosil models.Fosil
		if result.MatchedCount == 1 {
			err := fosilCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedFosil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.FosilResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedFosil}})
	}
}

// DeleteFosil godoc
// @Summary Delete a fosil by ID
// @Description Delete a fosil by ID
// @Tags Fosil
// @Param fosilId path string true "ID of the fosil to delete"
// @Produce json
// @Success 200 {object} responses.FosilResponse
// @Failure 404 {object} responses.FosilResponse
// @Failure 500 {object} responses.FosilResponse
// @Router /fosil/{fosilId} [delete]
func DeleteFosil() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		fosilId := c.Param("fosilId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(fosilId)

		result, err := fosilCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.FosilResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Fosil with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.FosilResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Fosil successfully deleted!"}},
		)
	}
}

// GetAllFosils godoc
// @Summary Get all fosils
// @Description Retrieve all fosils from the database
// @Tags Fosil
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.FosilResponse
// @Failure 500 {object} responses.FosilResponse
// @Router /fosils [get]
func GetAllFosils() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var fosils []models.Fosil
		defer cancel()

		results, err := fosilCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleFosil models.Fosil
			if err = results.Decode(&singleFosil); err != nil {
				c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			fosils = append(fosils, singleFosil)
		}

		c.JSON(http.StatusOK,
			responses.FosilResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": fosils}},
		)
	}
}

// ExportFosilToExcel export data fosil to excel.
// @Summary Export data fosil to excel
// @Description Get data fosil from MongoDB and export to excel
// @Tags Fosil
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Data fosil exported to excel successfully"
// @Failure 500 {object} responses.FosilResponse
// @Router /fosils/export [get]
func ExportFosilToExcel() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		file := xlsx.NewFile()
		sheet, err := file.AddSheet("Data Fosil")
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		rows, err := fosilCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
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
			var fosil models.Fosil
			if err := rows.Decode(&fosil); err != nil {
				c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			row := sheet.AddRow()
			row.AddCell().Value = fosil.Nomer.No_Reg
			row.AddCell().Value = fosil.Nomer.No_Inv
			row.AddCell().Value = fosil.Nomer.No_Awal
			row.AddCell().Value = fosil.Barang_Milik_Negara.Kode_Bmn
			row.AddCell().Value = fosil.Barang_Milik_Negara.Nup_Bmn
			row.AddCell().Value = fosil.Barang_Milik_Negara.Merk_Bmn
			row.AddCell().Value = fosil.Determinator
			row.AddCell().Value = fosil.Peta.Nama_Peta
			row.AddCell().Value = fosil.Peta.Skala_Peta
			row.AddCell().Value = fosil.Peta.Koleksi_Peta
			row.AddCell().Value = fosil.Peta.Lembar_Peta
			row.AddCell().Value = fosil.Cara_Perolehan
			row.AddCell().Value = fosil.Umur
			row.AddCell().Value = fosil.Nama_Satuan
			row.AddCell().Value = fosil.Kondisi
			row.AddCell().Value = fosil.Dalam_Negeri.Nama_Provinsi
			row.AddCell().Value = fosil.Dalam_Negeri.Nama_Kabupaten
			row.AddCell().Value = fosil.Luar_Negeri.Keterangan_LN
			row.AddCell().Value = fosil.Koleksi.Nama_Koleksi
			row.AddCell().Value = fosil.Koleksi.Jenis_Koleksi
			row.AddCell().Value = fosil.Koleksi.Sub_Jenis_Koleksi
			row.AddCell().Value = fosil.Koleksi.Kode_Jenis_Koleksi
			row.AddCell().Value = fosil.Koleksi.Deskripsi_Koleksi
			row.AddCell().Value = fosil.Koleksi.Kelompok_Koleksi
			row.AddCell().Value = fosil.Lokasi_Storage.Ruang_Storage
			row.AddCell().Value = fosil.Lokasi_Storage.Lantai
			row.AddCell().Value = fosil.Lokasi_Storage.Lajur
			row.AddCell().Value = fosil.Lokasi_Storage.Lemari
			row.AddCell().Value = fosil.Lokasi_Storage.Laci
			row.AddCell().Value = fosil.Lokasi_Storage.Slot
			row.AddCell().Value = fosil.Lokasi_Non_Storage.Nama_Non_Storage
			row.AddCell().Value = fosil.Nama_Formasi
			row.AddCell().Value = fosil.Keterangan
			row.AddCell().Value = fosil.Pulau
			row.AddCell().Value = fosil.Alamat_Lengkap
			row.AddCell().Value = fosil.Koordinat_X
			row.AddCell().Value = fosil.Koordinat_Y
			row.AddCell().Value = fosil.Koordinat_Z
			row.AddCell().Value = fosil.Tahun_Perolehan
			row.AddCell().Value = fosil.Kolektor
			row.AddCell().Value = fosil.Publikasi
			row.AddCell().Value = fosil.Kepemilikan_Awal
			row.AddCell().Value = fosil.URL
			row.AddCell().Value = fosil.Nilai_Perolehan
			row.AddCell().Value = fosil.Nilai_Buku
			row.AddCell().Value = fosil.Gambar_1
			row.AddCell().Value = fosil.Gambar_2
			row.AddCell().Value = fosil.Gambar_3
		}
		filename := "data_fosil.xlsx"
		err = file.Save(filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FosilResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")
		c.File(filename)
	}
}

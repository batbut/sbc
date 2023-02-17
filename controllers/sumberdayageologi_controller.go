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

var sumberdayageologiCollection *mongo.Collection = configs.GetCollection(configs.DB, "sumberdayageologi")
var validate_sumberdayageologi = validator.New()

// CreateSumberDayaGeologi godoc
// @Summary Create a new Sumber Daya Geologi
// @Description Create a new Sumber Daya Geologi with the input payload
// @Tags Sumber Daya Geologi
// @Accept  json
// @Produce  json
// @Param sumberdayageologi body models.SumberDayaGeologi true "The Sumber Daya Geologi to create"
// @Success 200 {object} responses.SumberDayaGeologiResponse
// @Failure 400 {object} responses.SumberDayaGeologiResponse
// @Router /sumberdayageologi [post]
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
			Barang_Milik_Negara: struct {
				Kode_Bmn string "bson:\"Kode_Bmn\" json:\"Kode_Bmn\" validate:\"required\""
				Nup_Bmn  string "bson:\"Nup_Bmn\" json:\"Nup_Bmn\" validate:\"required\""
				Merk_Bmn string "bson:\"Merk_Bmn\" json:\"Merk_Bmn\" validate:\"required\""
			}{
				Kode_Bmn: sumberdayageologi.Barang_Milik_Negara.Kode_Bmn,
				Nup_Bmn:  sumberdayageologi.Barang_Milik_Negara.Nup_Bmn,
				Merk_Bmn: sumberdayageologi.Barang_Milik_Negara.Merk_Bmn,
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
			Dalam_Negeri: struct {
				Nama_Provinsi  string "bson:\"Nama_Provinsi\" json:\"Nama_Provinsi\" validate:\"required\""
				Nama_Kabupaten string "bson:\"Nama_Kabupaten\" json:\"Nama_Kabupaten\" validate:\"required\""
			}{
				Nama_Provinsi:  sumberdayageologi.Dalam_Negeri.Nama_Provinsi,
				Nama_Kabupaten: sumberdayageologi.Dalam_Negeri.Nama_Kabupaten,
			},
			Luar_Negeri: struct {
				Keterangan_LN string "bson:\"Keterangan_LN\" json:\"Keterangan_LN\" validate:\"required\""
			}{
				Keterangan_LN: sumberdayageologi.Luar_Negeri.Keterangan_LN,
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

// GetSumberDayaGeologi godoc
// @Summary Get Sumber Daya Geologi By ID
// @Description Get a Sumber Daya Geologi By its ID
// @Tags Sumber Daya Geologi
// @ID get-sumberdayageologi-by-id
// @Produce json
// @Param sumberdayageologiId path string true "Sumber Daya Geologi ID"
// @Success 200 {object} responses.SumberDayaGeologiResponse
// @Failure 500 {object} responses.SumberDayaGeologiResponse
// @Router /sumberdayageologi/{sumberdayageologiId} [get]
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

// EditSumberDayaGeologi edits an existing Sumber Daya Geologi.
// @Summary Edit an existing Sumber Daya Geologi
// @Description Edit an existing Sumber Daya Geologi
// @Tags Sumber Daya Geologi
// @Accept json
// @Produce json
// @Param sumberdayageologiId path string true "ID of the Sumber Daya Geologi to edit"
// @Param body body models.SumberDayaGeologi true "Sumber Daya Geologi object that needs to be edited"
// @Success 200 {object} responses.SumberDayaGeologiResponse
// @Failure 400 {object} responses.SumberDayaGeologiResponse
// @Router /sumberdayageologi/{sumberdayageologiId} [put]
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
			"Barang_Milik_Negara": bson.M{
				"Kode_Bmn": sumberdayageologi.Barang_Milik_Negara.Kode_Bmn,
				"Nup_Bmn":  sumberdayageologi.Barang_Milik_Negara.Nup_Bmn,
				"Merk_Bmn": sumberdayageologi.Barang_Milik_Negara.Merk_Bmn,
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
			"Dalam_Negeri": bson.M{
				"Nama_Provinsi":  sumberdayageologi.Dalam_Negeri.Nama_Provinsi,
				"Nama_Kabupaten": sumberdayageologi.Dalam_Negeri.Nama_Kabupaten,
			},
			"Luar_Negeri": bson.M{
				"Keterangan_LN": sumberdayageologi.Luar_Negeri.Keterangan_LN,
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

// DeleteSumberDayaGeologi godoc
// @Summary Delete a Sumber Daya Geologi by ID
// @Description Delete a Sumber Daya Geologi by ID
// @Tags Sumber Daya Geologi
// @Param sumberdayageologiId path string true "ID of the Sumber Daya Geologi to delete"
// @Produce json
// @Success 200 {object} responses.SumberDayaGeologiResponse
// @Failure 404 {object} responses.SumberDayaGeologiResponse
// @Failure 500 {object} responses.SumberDayaGeologiResponse
// @Router /sumberdayageologi/{sumberdayageologiId} [delete]
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

// GetAllSumberDayaGeologis godoc
// @Summary Get all Sumber Daya Geologi
// @Description Retrieve all Sumber Daya Geologi from the database
// @Tags Sumber Daya Geologi
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.SumberDayaGeologiResponse
// @Failure 500 {object} responses.SumberDayaGeologiResponse
// @Router /sumberdayageologis [get]
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

// ExportSumberDayaGeologiToExcel export data Sumber Daya Geologi to excel.
// @Summary Export data Sumber Daya Geologi to excel
// @Description Get data Sumber Daya Geologi from MongoDB and export to excel
// @Tags Sumber Daya Geologi
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Data Sumber Daya Geologi exported to excel successfully"
// @Failure 500 {object} responses.SumberDayaGeologiResponse
// @Router /sumberdayageologis/export [get]
func ExportSumberDayaGeologiToExcel() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		file := xlsx.NewFile()
		sheet, err := file.AddSheet("Data Sumber Daya Geologi")
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		rows, err := sumberdayageologiCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
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
			var sumberdayageologi models.SumberDayaGeologi
			if err := rows.Decode(&sumberdayageologi); err != nil {
				c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			row := sheet.AddRow()
			row.AddCell().Value = sumberdayageologi.Nomer.No_Reg
			row.AddCell().Value = sumberdayageologi.Nomer.No_Inv
			row.AddCell().Value = sumberdayageologi.Nomer.No_Awal
			row.AddCell().Value = sumberdayageologi.Barang_Milik_Negara.Kode_Bmn
			row.AddCell().Value = sumberdayageologi.Barang_Milik_Negara.Nup_Bmn
			row.AddCell().Value = sumberdayageologi.Barang_Milik_Negara.Merk_Bmn
			row.AddCell().Value = sumberdayageologi.Determinator
			row.AddCell().Value = sumberdayageologi.Peta.Nama_Peta
			row.AddCell().Value = sumberdayageologi.Peta.Skala_Peta
			row.AddCell().Value = sumberdayageologi.Peta.Koleksi_Peta
			row.AddCell().Value = sumberdayageologi.Peta.Lembar_Peta
			row.AddCell().Value = sumberdayageologi.Cara_Perolehan
			row.AddCell().Value = sumberdayageologi.Umur
			row.AddCell().Value = sumberdayageologi.Nama_Satuan
			row.AddCell().Value = sumberdayageologi.Kondisi
			row.AddCell().Value = sumberdayageologi.Dalam_Negeri.Nama_Provinsi
			row.AddCell().Value = sumberdayageologi.Dalam_Negeri.Nama_Kabupaten
			row.AddCell().Value = sumberdayageologi.Luar_Negeri.Keterangan_LN
			row.AddCell().Value = sumberdayageologi.Koleksi.Nama_Koleksi
			row.AddCell().Value = sumberdayageologi.Koleksi.Jenis_Koleksi
			row.AddCell().Value = sumberdayageologi.Koleksi.Sub_Jenis_Koleksi
			row.AddCell().Value = sumberdayageologi.Koleksi.Kode_Jenis_Koleksi
			row.AddCell().Value = sumberdayageologi.Koleksi.Deskripsi_Koleksi
			row.AddCell().Value = sumberdayageologi.Koleksi.Kelompok_Koleksi
			row.AddCell().Value = sumberdayageologi.Lokasi_Storage.Ruang_Storage
			row.AddCell().Value = sumberdayageologi.Lokasi_Storage.Lantai
			row.AddCell().Value = sumberdayageologi.Lokasi_Storage.Lajur
			row.AddCell().Value = sumberdayageologi.Lokasi_Storage.Lemari
			row.AddCell().Value = sumberdayageologi.Lokasi_Storage.Laci
			row.AddCell().Value = sumberdayageologi.Lokasi_Storage.Slot
			row.AddCell().Value = sumberdayageologi.Lokasi_Non_Storage.Nama_Non_Storage
			row.AddCell().Value = sumberdayageologi.Nama_Formasi
			row.AddCell().Value = sumberdayageologi.Keterangan
			row.AddCell().Value = sumberdayageologi.Pulau
			row.AddCell().Value = sumberdayageologi.Alamat_Lengkap
			row.AddCell().Value = sumberdayageologi.Koordinat_X
			row.AddCell().Value = sumberdayageologi.Koordinat_Y
			row.AddCell().Value = sumberdayageologi.Koordinat_Z
			row.AddCell().Value = sumberdayageologi.Tahun_Perolehan
			row.AddCell().Value = sumberdayageologi.Kolektor
			row.AddCell().Value = sumberdayageologi.Publikasi
			row.AddCell().Value = sumberdayageologi.Kepemilikan_Awal
			row.AddCell().Value = sumberdayageologi.URL
			row.AddCell().Value = sumberdayageologi.Nilai_Perolehan
			row.AddCell().Value = sumberdayageologi.Nilai_Buku
			row.AddCell().Value = sumberdayageologi.Gambar_1
			row.AddCell().Value = sumberdayageologi.Gambar_2
			row.AddCell().Value = sumberdayageologi.Gambar_3
		}
		filename := "data_sdg.xlsx"
		err = file.Save(filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SumberDayaGeologiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")
		c.File(filename)
	}
}

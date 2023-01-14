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
			Id:               primitive.NewObjectID(),
			NoRegister:       batuan.NoRegister,
			NoInventaris:     batuan.NoInventaris,
			KodeBmn:          batuan.KodeBmn,
			NupBmn:           batuan.NupBmn,
			MerkBmn:          batuan.MerkBmn,
			Satuan:           batuan.Satuan,
			KelompokKoleksi:  batuan.KelompokKoleksi,
			JenisKoleksi:     batuan.JenisKoleksi,
			SubJenisKoleksi:  batuan.SubJenisKoleksi,
			KodeJenisKoleksi: batuan.KodeJenisKoleksi,
			RuangSimpan:      batuan.RuangSimpan,
			LokasiSimpan:     batuan.LokasiSimpan,
			Kondisi:          batuan.Kondisi,
			NamaKoleksi:      batuan.NamaKoleksi,
			Keterangan:       batuan.Keterangan,
			NamaFormasi:      batuan.NamaFormasi,
			LokasiTemuan:     batuan.LokasiTemuan,
			Koordinat:        batuan.Koordinat,
			Pulau:            batuan.Pulau,
			Peta:             batuan.Peta,
			LembarPeta:       batuan.LembarPeta,
			Skala:            batuan.Skala,
			CaraPerolehan:    batuan.CaraPerolehan,
			TahunPerolehan:   batuan.TahunPerolehan,
			Kolektor:         batuan.Kolektor,
			Kepemilikan:      batuan.Kepemilikan,
			Operator:         batuan.Operator,
			TanggalDicatat:   batuan.TanggalDicatat,
			NilaiPerolehan:   batuan.NilaiPerolehan,
			NilaiBuku:        batuan.NilaiBuku,
			Foto:             batuan.Foto,
			Foto2:            batuan.Foto2,
			Foto3:            batuan.Foto3,
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

		update := bson.M{"no_register": batuan.NoRegister, "no_inventaris": batuan.NoInventaris, "kode_bmn": batuan.KodeBmn, "nup_bmn": batuan.NupBmn, "merk_bmn": batuan.MerkBmn, "satuan": batuan.Satuan, "kelompok_koleksi": batuan.KelompokKoleksi, "jenis_koleksi": batuan.JenisKoleksi, "sub_jenis_koleksi": batuan.SubJenisKoleksi, "kode_jenis_koleksi": batuan.KodeJenisKoleksi, "ruang_simpan": batuan.RuangSimpan, "lokasi_simpan": batuan.LokasiSimpan, "kondisi": batuan.Kondisi, "nama_koleksi": batuan.NamaKoleksi, "keterangan": batuan.Keterangan, "nama_formasi": batuan.NamaFormasi, "lokasi_temuan": batuan.LokasiTemuan, "koordinat": batuan.Koordinat, "pulau": batuan.Pulau, "peta": batuan.Peta, "lembar_peta": batuan.LembarPeta, "skala": batuan.Skala, "cara_perolehan": batuan.CaraPerolehan, "tahun_perolehan": batuan.TahunPerolehan, "kolektor": batuan.Kolektor, "kepemilikan": batuan.Kepemilikan, "operator": batuan.Operator, "tanggal_dicatat": batuan.TanggalDicatat, "nilai_perolehan": batuan.NilaiPerolehan, "nilai_buku": batuan.NilaiBuku, "foto": batuan.Foto, "foto_2": batuan.Foto2, "foto_3": batuan.Foto3}
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

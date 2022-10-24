package productcontroller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/roneyhendrix/Golang-rest-sederhana/helper"
	"github.com/roneyhendrix/Golang-rest-sederhana/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func ShowAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if err := models.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, products)
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product tidak ditemukan")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusCreated, product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	//update product
	if models.DB.Model(&product).Where("id = ?", id).Updates(product).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak dapat mengupdate product")
		return
	}
	product.Id = id
	ResponseJson(w, http.StatusOK, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	var product models.Product
	if models.DB.Delete(&product, input["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak dapat menghapus product")
		return
	}
	response := map[string]string{"message": "Product berhasil dihapus"}
	ResponseJson(w, http.StatusOK, response)
}

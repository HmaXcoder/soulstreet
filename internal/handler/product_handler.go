package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"soulstreet/internal/model"
	"soulstreet/internal/service"
	"soulstreet/pkg/json"
	"strconv"

	"github.com/google/uuid"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler (productService service.ProductService) *ProductHandler {
	return &ProductHandler {
		productService: productService,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	
	r.ParseMultipartForm(10 << 20)
	name := r.FormValue("name")
	priceStr := r.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		json.SendJsonError(w, http.StatusBadRequest, errors.New("Erro ao parsear 'price'"))
		return
	}
	files := r.MultipartForm.File["images"]
	var imagesPaths []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			json.SendJsonError(w, http.StatusInternalServerError, errors.New("Erro ao processar as imagens"))
			return
		}
		defer file.Close()
		newFileName := uuid.New().String() + filepath.Ext(fileHeader.Filename)
		savePath := "uploads/" + newFileName
		
		outFile, err := os.Create(savePath)
		if err != nil {
			json.SendJsonError(w, http.StatusInternalServerError, errors.New("Erro ao salvar imagem"))
			return
		}
		defer outFile.Close()
		
		_, err = io.Copy(outFile, file)
		if err != nil {
			json.SendJsonError(w, http.StatusInternalServerError, errors.New("Erro ao escrever imagem"))
			return
		}
		imagesPaths = append(imagesPaths, savePath)
	}
	product := model.Product {
		Name: name,
		Price: float32(price),
		Images: imagesPaths,
	}
	err = h.productService.CreateProduct(&product)
	if err != nil {
		json.SendJsonError(w, http.StatusInternalServerError, errors.New("Erro ao salvar produto"))
		return
	}
	json.SendJson(w, http.StatusCreated, product)
}


func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
	json.SendJsonError(w, http.StatusBadRequest, errors.New("Query 'id' em branco"))
	return
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		json.SendJsonError(w, http.StatusBadRequest, errors.New("Erro ao parsear 'id'"))
		return
	}
	product, err := h.productService.GetProductByID(int(id))
	if err != nil {
		json.SendJsonError(w, http.StatusBadRequest, fmt.Errorf("error: %v", err))
		return
	}
	json.SendJson(w, http.StatusOK, product)

}


func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAll()
	if err != nil {
		json.SendJsonError(w, http.StatusInternalServerError, err)
		return
	}
	json.SendJson(w, http.StatusOK, products)
}
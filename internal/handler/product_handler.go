package handler

import (
	"errors"
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
	
	if r.Method != http.MethodPost {
		json.SendJsonError(w, http.StatusMethodNotAllowed, errors.New("Metodo não permitido!"))
		return
	}
	
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
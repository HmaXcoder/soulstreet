package repository

import (
	"database/sql"
	"fmt"
	"soulstreet/internal/model"
)

type ProductRepository interface {
	Create(product *model.Product) error
	GetByID(id int) (*model.Product, error)
	GetAll() ([]model.Product, error)
}

type productRepositoryDB struct {
	db *sql.DB
}


func NewProductRepositoryDB(db *sql.DB) ProductRepository {
	return &productRepositoryDB{db: db}
}

func (r *productRepositoryDB) Create(product *model.Product) error {
	query := "INSERT INTO products (name, price, images, avaliable) VALUES (?,?,?,?)"
	_, err := r.db.Exec(query, product.Name, product.Price, product.Images, product.IsAvaliable)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepositoryDB) GetByID(id int) (*model.Product, error) {
	var product model.Product
	query := "SELECT id, name, price, images, avaliable FROM products WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Images, &product.IsAvaliable)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("produto não encontrado")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepositoryDB) GetAll() ([]model.Product, error) {
	var products []model.Product
	rows, err := r.db.Query("SELECT id, name, price, images, avaliable FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Images, &product.IsAvaliable); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
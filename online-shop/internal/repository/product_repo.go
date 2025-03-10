package repository

import (
	"database/sql"
	"online-shop/internal/models"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (r *ProductRepo) GetProductsByName(name string) ([]models.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, description, price, category FROM products WHERE name ILIKE '%' || $1 || '%'", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

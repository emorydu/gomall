package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:products_categories"`
}

func (p Product) TableName() string {
	return "products"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

func (p ProductQuery) GetByID(productID int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productID).Error
	return
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name LIKE ? OR description LIKE ?",
		"%"+q+"%", "%"+q+"%").Error
	return
}

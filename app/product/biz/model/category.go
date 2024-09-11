package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"products" gorm:"many2many:products_categories"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}

func (c CategoryQuery) GetProductsByCategoryName(name string) (categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&categories).Error
	return
}

package service

import (
	"context"

	"github.com/emory/gomall/app/product/biz/dal/mysql"
	"github.com/emory/gomall/app/product/biz/model"
	product "github.com/emory/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)

	resp = &product.ListProductsResp{}
	for _, cc := range c {
		for _, v := range cc.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
			})
		}
	}

	return resp, err
}

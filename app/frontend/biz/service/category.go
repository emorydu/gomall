package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/emory/gomall/app/frontend/infra/rpc"
	"github.com/emory/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	category "github.com/emory/gomall/app/frontend/hertz_gen/frontend/category"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{
		CategoryName: req.Category,
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title": "Category",
		"Items": p.Products,
	}, nil
}

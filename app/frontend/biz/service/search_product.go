package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/emory/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/emory/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	product "github.com/emory/gomall/app/frontend/hertz_gen/frontend/product"
)

type SearchProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductService(Context context.Context, RequestContext *app.RequestContext) *SearchProductService {
	return &SearchProductService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductService) Run(req *product.SearchProductReq) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Items": products.Results,
		"Q":     req.Q,
	}, nil
}

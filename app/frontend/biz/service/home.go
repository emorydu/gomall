package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/emory/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/emory/gomall/app/frontend/infra/rpc"
	"github.com/emory/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title": "Hot Sale",
		"Items": products,
	}, nil
}

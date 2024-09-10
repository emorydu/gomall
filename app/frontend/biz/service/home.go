package service

import (
	"context"
	"github.com/emory/gomall/app/frontend/hertz_gen/frontend/common"

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
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "T-Shirt-1", "Price": 100, "Picture": "./static/image/t-shirt-1.jpeg"},
		{"Name": "T-Shirt-2", "Price": 110, "Picture": "./static/image/t-shirt-1.jpeg"},
		{"Name": "T-Shirt-3", "Price": 120, "Picture": "./static/image/t-shirt-2.jpeg"},
		{"Name": "T-Shirt-4", "Price": 130, "Picture": "./static/image/notebook.jpeg"},
		{"Name": "T-Shirt-5", "Price": 140, "Picture": "./static/image/t-shirt-1.jpeg"},
		{"Name": "T-Shirt-6", "Price": 150, "Picture": "./static/image/t-shirt.jpeg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items

	return resp, nil
}

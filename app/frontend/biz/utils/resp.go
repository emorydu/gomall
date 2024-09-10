package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/emory/gomall/app/frontend/utils"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WrapResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	content["UserID"] = ctx.Value(utils.GetUserIdFromCtx(ctx))
	return content
}

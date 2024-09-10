package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) uint32 {
	userID := ctx.Value(SessionUserId)
	if userID == nil {
		return 0
	}
	return userID.(uint32)
}

package xuser

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"interface/utility/utils/middleware"
)

func Uid(ctx context.Context) uint64 {
	return g.RequestFromCtx(ctx).Get(middleware.Uid).Uint64()
}

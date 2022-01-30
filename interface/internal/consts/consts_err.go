package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ErrService              = gerror.NewCode(gcode.CodeInternalError, "系统错误")
	ErrIllegal              = gerror.NewCode(gcode.CodeOperationFailed, "非法操作")
	ErrAuthNotEnough        = gerror.NewCode(gcode.CodeOperationFailed, "暂无当前操作权限")
	ErrAuth                 = gerror.NewCode(gcode.CodeNotAuthorized, "未认证")
	ErrLogin                = gerror.NewCode(gcode.CodeValidationFailed, "用户名或密码错误")
	ErrUnameAlreadyExist    = gerror.NewCode(gcode.CodeValidationFailed, "用户名已存在")
	ErrDataNotFound         = gerror.NewCode(gcode.CodeNotFound, "数据不存在")
	ErrClose                = gerror.NewCode(gcode.CodeNotSupported, "暂不支持")
	ErrWriterEx1IsEmpty     = gerror.NewCode(gcode.CodeValidationFailed, "作者出生时间不能为空")
	ErrWriterBirthdayFormat = gerror.NewCode(gcode.CodeInvalidParameter, "解析参数失败了")
	ErrAddBookWithoutYear   = gerror.NewCode(gcode.CodeValidationFailed, "发布年份不能为空")
)

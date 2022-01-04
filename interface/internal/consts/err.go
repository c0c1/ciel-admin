package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ErrService           = gerror.NewCode(gcode.CodeInternalError, "Internal Error")
	ErrIllegal           = gerror.NewCode(gcode.CodeOperationFailed, "Operation Failed")
	ErrAuthNotEnough     = gerror.NewCode(gcode.CodeValidationFailed, "Validation Failed")
	ErrAuth              = gerror.NewCode(gcode.CodeNotAuthorized, "Not Authorized")
	ErrLogin             = gerror.NewCode(gcode.CodeValidationFailed, "Uname or Pwd is error")
	ErrUnameAlreadyExist = gerror.NewCode(gcode.CodeValidationFailed, "Uname Already Exist")
	ErrDataNotFound      = gerror.NewCode(gcode.CodeNotFound, "Not Found")
	ErrClose             = gerror.NewCode(gcode.CodeNotSupported, "Not Supported")
	ErrOldPwdIsError     = gerror.NewCode(gcode.CodeValidationFailed, "Old Pwd is Error")
	ErrOldNewPwdIsSame   = gerror.NewCode(gcode.CodeValidationFailed, "Pwd Can't be Same")
)

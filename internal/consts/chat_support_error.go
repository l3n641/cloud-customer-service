package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ChatSupportNoAuth = gerror.NewCode(gcode.CodeNotAuthorized, "没有权限")
)

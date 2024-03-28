package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ClientNotAuth     = gerror.NewCode(gcode.CodeNotAuthorized, "Not Auth")
	ClientNotTicket   = gerror.NewCode(gcode.CodeInvalidConfiguration, "Invalid Configuration")
	NotHasChatSupport = gerror.NewCode(gcode.CodeInvalidConfiguration, "Not Has Chat Support")
)

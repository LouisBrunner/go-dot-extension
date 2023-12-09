package gdextension

import (
	"fmt"
	"log"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

func (me *extension) Logf(level LogLevel, format string, args ...interface{}) {
	me.LogDetailedf(level, "go-dot-extension", "internal", "internal", 0, false, format, args...)
}

func (me *extension) LogDetailedf(level LogLevel, description, function, file string, line int32, notifyEditor bool, format string, args ...interface{}) {
	if level > me.logLevel {
		return
	}

	msg := fmt.Sprintf(format, args...)

	lineC := int(line)
	notifyEditorC := gdc.Bool(0)
	if notifyEditor {
		notifyEditorC = gdc.Bool(1)
	}

	editor := ""
	if notifyEditor {
		editor = " {EDITOR}"
	}
	log.Printf("[%s] %s:%d @ %s(%s): %s%s", level, file, line, function, description, msg, editor)
	switch level {
	case LogLevelError:
		me.iface.PrintErrorWithMessage(description, msg, function, file, lineC, notifyEditorC)
	case LogLevelWarning:
		me.iface.PrintWarningWithMessage(description, msg, function, file, lineC, notifyEditorC)
	case LogLevelInfo:
		gdapi.Utilities.Print(gdapi.NewVariantFrom(gdapi.StringFromStr(fmt.Sprintf("%s: %s", description, msg))))
	case LogLevelDebug:
	default:
		log.Printf("unknown log level %q for %q", level, msg)
		me.iface.PrintWarningWithMessage(description, fmt.Sprintf("UNKNOWN LEVEL %q: %s", level, msg), function, file, lineC, notifyEditorC)
	}
}

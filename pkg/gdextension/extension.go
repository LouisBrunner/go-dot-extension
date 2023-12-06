package gdextension

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type extension struct {
	iface     gdc.Interface
	pLibrary  gdc.ClassLibraryPtr
	initLevel gdc.InitializationLevel

	toRegister []ClassConstructor
	registered map[string]*classEntry

	logLevel LogLevel
}

func New(pGetProcAddr gdc.InterfaceGetProcAddress, pLibrary gdc.ClassLibraryPtr, logLevel LogLevel) (Extension, error) {
	iface, err := gdc.NewInterface(pGetProcAddr)
	if err != nil {
		return nil, fmt.Errorf("could not get GDExtension class: %s", err.Error())
	}
	ext := &extension{
		iface:     iface,
		pLibrary:  pLibrary,
		initLevel: gdc.InitializationScene,
		logLevel:  logLevel,
	}
	return ext, nil
}

func (me *extension) registerClass(className, parentName string, def gdc.ClassCreationInfo) error {
	// return me.callInterfaceClassdbRegisterExtensionClass(
	//
	//	gdc.ClassLibraryPtr(me.pLibrary),
	//	className,
	//	parentName,
	//	UnsafeConvert[*gdc.ClassCreationInfo](&def),
	//
	// )
	return nil
}

func (me *extension) unregisterClass(className string) error {
	// return me.callInterfaceClassdbUnregisterExtensionClass(
	// 	gdc.ClassLibraryPtr(me.pLibrary),
	// 	className,
	// )
	return nil
}

func (me *extension) Register(classes ...ClassConstructor) {
	me.toRegister = append(me.toRegister, classes...)
}

func (me *extension) SetInitializationLevel(level gdc.InitializationLevel) {
	lowestAllowed := gdc.InitializationLevel(gdc.InitializationScene)
	if level < lowestAllowed {
		me.Logf(LogLevelWarning, "setting initialization level to %v, which is lower than %v, classes will not be added", level, lowestAllowed)
	}
	me.initLevel = level
}

func (me *extension) Initialize(rInitialization *gdc.InitializationRaw, init gdc.InitializationInitializeFn, fini gdc.InitializationDeinitializeFn) gdc.Bool {
	me.registered = make(map[string]*classEntry, len(me.toRegister))
	for _, constructor := range me.toRegister {
		instance := constructor()
		me.Logf(LogLevelDebug, "adding class %q to tracked classes", instance.ClassName())
		me.registered[instance.ClassName()] = &classEntry{
			ext:         me,
			class:       instance,
			constructor: constructor,
			instances:   make(map[uint64]struct{}),
		}
	}
	rInit := gdc.Initialization{
		MinimumInitializationLevel: me.initLevel,
		Userdata:                   store(me),
		Initialize:                 init,
		Deinitialize:               fini,
	}
	*rInitialization = rInit.ToRaw()
	return gdc.Bool(1)
}

func (me *extension) OnInit(level gdc.InitializationLevel) error {
	if level != gdc.InitializationLevel(gdc.InitializationScene) {
		return nil
	}
	for name, entry := range me.registered {
		me.Logf(LogLevelDebug, "registering class %q", name)
		err := me.registerClass(entry.class.ClassName(), entry.class.ParentClassName(), gdc.ClassCreationInfo{
			ClassUserdata: store(entry),
			// CreateInstanceFunc: trampolineClassCreateInstance,
			// FreeInstanceFunc:   trampolineClassFreeInstance,
		})
		if err != nil {
			return fmt.Errorf("could not register class %q: %s", name, err.Error())
		}
	}
	return nil
}

func (me *extension) OnFini(level gdc.InitializationLevel) error {
	if level != gdc.InitializationLevel(gdc.InitializationScene) {
		return nil
	}
	for name := range me.registered {
		me.Logf(LogLevelDebug, "unregistering class %q", name)
		err := me.unregisterClass(name)
		if err != nil {
			return fmt.Errorf("could not unregister class %q: %s", name, err.Error())
		}
	}
	return nil
}

func RestoreFromC(pUserData unsafe.Pointer) (Extension, error) {
	return restore[*extension](pUserData)
}

func (me *extension) Logf(level LogLevel, format string, args ...interface{}) {
	me.LogDetailedf(level, "go-dot-extension", "internal", "internal", 0, false, format, args...)
}

func (me *extension) LogDetailedf(level LogLevel, description, function, file string, line int32, notifyEditor bool, format string, args ...interface{}) {
	if level > me.logLevel {
		return
	}

	msg := fmt.Sprintf(format, args...)

	editor := ""
	if notifyEditor {
		editor = " {EDITOR}"
	}
	log.Printf("[%s] %s:%d @ %s(%s): %s%s", level, file, line, function, description, msg, editor)
	var err error
	switch level {
	case LogLevelError:
		// err = me.callInterfacePrintErrorWithMessage(description, msg, function, file, line, notifyEditor)
	case LogLevelWarning:
		// err = me.callInterfacePrintWarningWithMessage(description, msg, function, file, line, notifyEditor)
	case LogLevelInfo:
		// err = me.callPrint(fmt.Sprintf("%s: %s", description, msg))
	case LogLevelDebug:
	default:
		log.Printf("unknown log level %q for %q", level, msg)
		// err = me.callInterfacePrintWarningWithMessage(description, fmt.Sprintf("UNKNOWN LEVEL %q: %s", level, msg), function, file, line, notifyEditor)
	}
	if err != nil {
		log.Printf("could not log message to editor: %v", err.Error())
	}
}

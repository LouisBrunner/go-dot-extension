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

	gdc.Callbacks.SetInitializationInitializeHandler(func(userdata unsafe.Pointer, pLevel gdc.InitializationLevel) {
		ext.Logf(LogLevelDebug, "initializing module (level=%v)", pLevel)
		err := ext.onInit(pLevel)
		if err != nil {
			ext.Logf(LogLevelError, "error: %s", err.Error())
		}
	})
	gdc.Callbacks.SetInitializationDeinitializeHandler(func(userdata unsafe.Pointer, pLevel gdc.InitializationLevel) {
		ext.Logf(LogLevelDebug, "uninitializing module (level=%v)", pLevel)
		err := ext.onFini(pLevel)
		if err != nil {
			ext.Logf(LogLevelError, "error: %s", err.Error())
		}
	})
	gdc.Callbacks.SetClassCreationInfoCreateInstanceFuncHandler(func(pUserdata unsafe.Pointer) gdc.ObjectPtr {
		class, err := restore[*classEntry](pUserdata)
		if err != nil {
			ext.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return nil
		}
		return ext.createClass(class)
	})
	gdc.Callbacks.SetClassCreationInfoFreeInstanceFuncHandler(func(pUserdata unsafe.Pointer, pInstance gdc.ClassInstancePtr) {
		class, err := restore[*classEntry](pUserdata)
		if err != nil {
			ext.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return
		}
		ext.freeClass(class, pInstance)
	})

	return ext, nil
}

func (me *extension) createClass(class *classEntry) gdc.ObjectPtr {
	parent, clean := me.makeStringName(class.class.ParentClassName())
	defer clean()
	name, clean := me.makeStringName(class.class.ClassName())
	defer clean()

	obj := me.iface.ClassdbConstructObject(gdc.ConstStringNamePtr(parent))
	id := me.iface.ObjectGetInstanceId(gdc.ConstObjectPtr(obj))
	me.iface.ObjectSetInstance(obj, gdc.ConstStringNamePtr(name), gdc.ClassInstancePtr(&id))
	class.addInstance(uint64(id))
	return obj
}

func (me *extension) freeClass(class *classEntry, pInstance gdc.ClassInstancePtr) {
	id := *(*uint64)(pInstance)
	class.deleteInstance(uint64(id))
}

func (me *extension) registerClass(className, parentName string, def gdc.ClassCreationInfo) error {
	class, clean := me.makeStringName(className)
	defer clean()
	parent, clean := me.makeStringName(parentName)
	defer clean()
	me.iface.ClassdbRegisterExtensionClass(me.pLibrary, gdc.ConstStringNamePtr(class), gdc.ConstStringNamePtr(parent), &def)
	return nil
}

func (me *extension) unregisterClass(className string) error {
	class, clean := me.makeStringName(className)
	defer clean()
	me.iface.ClassdbUnregisterExtensionClass(me.pLibrary, gdc.ConstStringNamePtr(class))
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
	*rInitialization, _ = rInit.ToRaw()
	return gdc.Bool(1)
}

func (me *extension) onInit(level gdc.InitializationLevel) error {
	if level != gdc.InitializationLevel(gdc.InitializationScene) {
		return nil
	}
	for name, entry := range me.registered {
		me.Logf(LogLevelDebug, "registering class %q", name)
		err := me.registerClass(entry.class.ClassName(), entry.class.ParentClassName(), gdc.ClassCreationInfo{
			ClassUserdata:      store(entry),
			CreateInstanceFunc: gdc.Callbacks.GetClassCreationInfoCreateInstanceFuncCallback(),
			FreeInstanceFunc:   gdc.Callbacks.GetClassCreationInfoFreeInstanceFuncCallback(),
		})
		if err != nil {
			return fmt.Errorf("could not register class %q: %s", name, err.Error())
		}
	}
	return nil
}

func (me *extension) onFini(level gdc.InitializationLevel) error {
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
		me.callPrint(fmt.Sprintf("%s: %s", description, msg))
	case LogLevelDebug:
	default:
		log.Printf("unknown log level %q for %q", level, msg)
		me.iface.PrintWarningWithMessage(description, fmt.Sprintf("UNKNOWN LEVEL %q: %s", level, msg), function, file, lineC, notifyEditorC)
	}
}

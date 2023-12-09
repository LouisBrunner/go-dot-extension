package gdextension

import (
	"fmt"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
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
	gdapi.Initialize(iface)

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

	ext.addClassCallbacks()
	return ext, nil
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
		me.Logf(LogLevelDebug, "adding class %T to tracked classes", instance)
		entry, err := me.prepareClass(constructor)
		if err != nil {
			me.Logf(LogLevelError, "could not prepare class %T: %s", instance, err.Error())
			return gdc.Bool(0)
		}
		me.registered[entry.name] = entry
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
	if level != gdc.InitializationScene {
		return nil
	}
	gdapi.InitializeGlobals(me.iface)
	for name, entry := range me.registered {
		me.Logf(LogLevelDebug, "registering class %q", name)
		me.registerClass(entry)
	}
	return nil
}

func (me *extension) onFini(level gdc.InitializationLevel) error {
	if level != gdc.InitializationScene {
		return nil
	}
	for name, entry := range me.registered {
		me.Logf(LogLevelDebug, "unregistering class %q", name)
		me.unregisterClass(entry)
	}
	return nil
}

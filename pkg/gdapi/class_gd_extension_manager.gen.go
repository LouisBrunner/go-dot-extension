// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForGDExtensionManagerList struct {
	fnLoadExtension       gdc.MethodBindPtr
	fnReloadExtension     gdc.MethodBindPtr
	fnUnloadExtension     gdc.MethodBindPtr
	fnIsExtensionLoaded   gdc.MethodBindPtr
	fnGetLoadedExtensions gdc.MethodBindPtr
	fnGetExtension        gdc.MethodBindPtr
}

var ptrsForGDExtensionManager ptrsForGDExtensionManagerList

func initGDExtensionManagerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GDExtensionManager")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("load_extension")
		defer methodName.Destroy()
		ptrsForGDExtensionManager.fnLoadExtension = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4024158731))
	}
	{
		methodName := StringNameFromStr("reload_extension")
		defer methodName.Destroy()
		ptrsForGDExtensionManager.fnReloadExtension = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4024158731))
	}
	{
		methodName := StringNameFromStr("unload_extension")
		defer methodName.Destroy()
		ptrsForGDExtensionManager.fnUnloadExtension = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4024158731))
	}
	{
		methodName := StringNameFromStr("is_extension_loaded")
		defer methodName.Destroy()
		ptrsForGDExtensionManager.fnIsExtensionLoaded = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("get_loaded_extensions")
		defer methodName.Destroy()
		ptrsForGDExtensionManager.fnGetLoadedExtensions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("get_extension")
		defer methodName.Destroy()
		ptrsForGDExtensionManager.fnGetExtension = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 49743343))
	}

}

type GDExtensionManager struct {
	Object
}

func (me *GDExtensionManager) BaseClass() string {
	return "GDExtensionManager"
}

func NewGDExtensionManager() *GDExtensionManager {
	str := StringNameFromStr("GDExtensionManager") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GDExtensionManager{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type GDExtensionManagerLoadStatus int

const (
	GDExtensionManagerLoadStatusLoadStatusOk            GDExtensionManagerLoadStatus = 0
	GDExtensionManagerLoadStatusLoadStatusFailed        GDExtensionManagerLoadStatus = 1
	GDExtensionManagerLoadStatusLoadStatusAlreadyLoaded GDExtensionManagerLoadStatus = 2
	GDExtensionManagerLoadStatusLoadStatusNotLoaded     GDExtensionManagerLoadStatus = 3
	GDExtensionManagerLoadStatusLoadStatusNeedsRestart  GDExtensionManagerLoadStatus = 4
)

func (me *GDExtensionManager) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GDExtensionManager) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GDExtensionManager) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GDExtensionManager) LoadExtension(path String) GDExtensionManagerLoadStatus {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GDExtensionManagerLoadStatus

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtensionManager.fnLoadExtension), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GDExtensionManager) ReloadExtension(path String) GDExtensionManagerLoadStatus {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GDExtensionManagerLoadStatus

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtensionManager.fnReloadExtension), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GDExtensionManager) UnloadExtension(path String) GDExtensionManagerLoadStatus {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GDExtensionManagerLoadStatus

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtensionManager.fnUnloadExtension), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GDExtensionManager) IsExtensionLoaded(path String) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtensionManager.fnIsExtensionLoaded), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GDExtensionManager) GetLoadedExtensions() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtensionManager.fnGetLoadedExtensions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GDExtensionManager) GetExtension(path String) GDExtension {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGDExtension()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtensionManager.fnGetExtension), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals

type GDExtensionManagerExtensionsReloadedSignalFn func()

func (me *GDExtensionManager) ConnectExtensionsReloaded(subs SignalSubscribers, fn GDExtensionManagerExtensionsReloadedSignalFn) {
	sig := StringNameFromStr("extensions_reloaded")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GDExtensionManager) DisconnectExtensionsReloaded(subs SignalSubscribers, fn GDExtensionManagerExtensionsReloadedSignalFn) {
	sig := StringNameFromStr("extensions_reloaded")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

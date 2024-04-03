// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  GDExtensionManagerLoadStatusLoadStatusOk GDExtensionManagerLoadStatus = 0
  GDExtensionManagerLoadStatusLoadStatusFailed GDExtensionManagerLoadStatus = 1
  GDExtensionManagerLoadStatusLoadStatusAlreadyLoaded GDExtensionManagerLoadStatus = 2
  GDExtensionManagerLoadStatusLoadStatusNotLoaded GDExtensionManagerLoadStatus = 3
  GDExtensionManagerLoadStatusLoadStatusNeedsRestart GDExtensionManagerLoadStatus = 4
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

func  (me *GDExtensionManager) LoadExtension(path String, ) GDExtensionManagerLoadStatus {
  classNameV := StringNameFromStr("GDExtensionManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4024158731) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  var ret GDExtensionManagerLoadStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GDExtensionManager) ReloadExtension(path String, ) GDExtensionManagerLoadStatus {
  classNameV := StringNameFromStr("GDExtensionManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reload_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4024158731) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  var ret GDExtensionManagerLoadStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GDExtensionManager) UnloadExtension(path String, ) GDExtensionManagerLoadStatus {
  classNameV := StringNameFromStr("GDExtensionManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unload_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4024158731) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  var ret GDExtensionManagerLoadStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GDExtensionManager) IsExtensionLoaded(path String, ) bool {
  classNameV := StringNameFromStr("GDExtensionManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_extension_loaded")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GDExtensionManager) GetLoadedExtensions() PackedStringArray {
  classNameV := StringNameFromStr("GDExtensionManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loaded_extensions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GDExtensionManager) GetExtension(path String, ) GDExtension {
  classNameV := StringNameFromStr("GDExtensionManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 49743343) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewGDExtension()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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

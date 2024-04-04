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

type JavaScriptBridge struct {
  Object
}

func (me *JavaScriptBridge) BaseClass() string {
  return "JavaScriptBridge"
}

func NewJavaScriptBridge() *JavaScriptBridge {
  str := StringNameFromStr("JavaScriptBridge") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &JavaScriptBridge{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *JavaScriptBridge) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JavaScriptBridge) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JavaScriptBridge) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *JavaScriptBridge) Eval(code String, use_global_execution_context bool, ) Variant {
  classNameV := StringNameFromStr("JavaScriptBridge")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("eval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 218087648) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{code.AsCTypePtr(), gdc.ConstTypePtr(&use_global_execution_context) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&use_global_execution_context)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JavaScriptBridge) GetInterface(interface_ String, ) JavaScriptObject {
  classNameV := StringNameFromStr("JavaScriptBridge")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1355533281) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewJavaScriptObject()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JavaScriptBridge) CreateCallback(callable Callable, ) JavaScriptObject {
  classNameV := StringNameFromStr("JavaScriptBridge")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 422818440) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewJavaScriptObject()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JavaScriptBridge) CreateObject(object String, varargs ...Variant) Variant {
  classNameV := StringNameFromStr("JavaScriptBridge")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3093893586) // FIXME: should cache?
  cargs := make([]gdc.ConstVariantPtr, 0, 1 + len(varargs))
  var0 := object.AsVariant()
  defer var0.Destroy()
  cargs = append(cargs, var0.AsCPtr())
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }
  ret := NewVariant()
  cerr := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return *ret
  }
  return *ret
}

func  (me *JavaScriptBridge) DownloadBuffer(buffer PackedByteArray, name String, mime String, )  {
  classNameV := StringNameFromStr("JavaScriptBridge")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("download_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3352272093) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), name.AsCTypePtr(), mime.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *JavaScriptBridge) PwaNeedsUpdate() bool {
  classNameV := StringNameFromStr("JavaScriptBridge")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pwa_needs_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *JavaScriptBridge) PwaUpdate() Error {
  classNameV := StringNameFromStr("JavaScriptBridge")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pwa_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *JavaScriptBridge) ForceFsSync()  {
  classNameV := StringNameFromStr("JavaScriptBridge")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_fs_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type JavaScriptBridgePwaUpdateAvailableSignalFn func()

func (me *JavaScriptBridge) ConnectPwaUpdateAvailable(subs SignalSubscribers, fn JavaScriptBridgePwaUpdateAvailableSignalFn) {
  sig := StringNameFromStr("pwa_update_available")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *JavaScriptBridge) DisconnectPwaUpdateAvailable(subs SignalSubscribers, fn JavaScriptBridgePwaUpdateAvailableSignalFn) {
  sig := StringNameFromStr("pwa_update_available")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

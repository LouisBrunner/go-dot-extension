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

type ptrsForHMACContextList struct {
  fnStart gdc.MethodBindPtr
  fnUpdate gdc.MethodBindPtr
  fnFinish gdc.MethodBindPtr
}

var ptrsForHMACContext ptrsForHMACContextList

func initHMACContextPtrs(iface gdc.Interface) {

  className := StringNameFromStr("HMACContext")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("start")
    defer methodName.Destroy()
    ptrsForHMACContext.fnStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3537364598))
  }
  {
    methodName := StringNameFromStr("update")
    defer methodName.Destroy()
    ptrsForHMACContext.fnUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
  }
  {
    methodName := StringNameFromStr("finish")
    defer methodName.Destroy()
    ptrsForHMACContext.fnFinish = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2115431945))
  }
}

type HMACContext struct {
  RefCounted
}

func (me *HMACContext) BaseClass() string {
  return "HMACContext"
}

func NewHMACContext() *HMACContext {
  str := StringNameFromStr("HMACContext") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &HMACContext{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *HMACContext) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HMACContext) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HMACContext) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *HMACContext) Start(hash_type HashingContextHashType, key PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hash_type) , key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&hash_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHMACContext.fnStart), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HMACContext) Update(data PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHMACContext.fnUpdate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HMACContext) Finish() PackedByteArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHMACContext.fnFinish), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

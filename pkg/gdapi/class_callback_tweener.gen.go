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

type ptrsForCallbackTweenerList struct {
  fnSetDelay gdc.MethodBindPtr
}

var ptrsForCallbackTweener ptrsForCallbackTweenerList

func initCallbackTweenerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CallbackTweener")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_delay")
    defer methodName.Destroy()
    ptrsForCallbackTweener.fnSetDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3008182292))
  }
}

type CallbackTweener struct {
  Tweener
}

func (me *CallbackTweener) BaseClass() string {
  return "CallbackTweener"
}

func NewCallbackTweener() *CallbackTweener {
  str := StringNameFromStr("CallbackTweener") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CallbackTweener{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CallbackTweener) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CallbackTweener) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CallbackTweener) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CallbackTweener) SetDelay(delay float64, ) CallbackTweener {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCallbackTweener()
  pinner.Pin(&delay)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCallbackTweener.fnSetDelay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

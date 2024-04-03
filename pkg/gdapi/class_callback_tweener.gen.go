// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("CallbackTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3008182292) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay), }
  ret := NewCallbackTweener()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CallbackTweener struct {
  obj gdc.ObjectPtr
}

func (me *CallbackTweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CallbackTweener) BaseClass() string {
  return "CallbackTweener"
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

func  (me *CallbackTweener) SetDelay(delay float32, ) CallbackTweener {
  classNameV := StringNameFromStr("CallbackTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3008182292) // FIXME: should cache?
  var ret CallbackTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals

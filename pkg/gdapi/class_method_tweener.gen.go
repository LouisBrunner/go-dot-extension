// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MethodTweener struct {
  obj gdc.ObjectPtr
}

func (me *MethodTweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MethodTweener) BaseClass() string {
  return "MethodTweener"
}



// Enums

func (me *MethodTweener) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MethodTweener) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MethodTweener) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MethodTweener) SetDelay(delay float32, ) MethodTweener {
  classNameV := StringNameFromStr("MethodTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 266477812) // FIXME: should cache?
  var ret MethodTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MethodTweener) SetTrans(trans TweenTransitionType, ) MethodTweener {
  classNameV := StringNameFromStr("MethodTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_trans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740975367) // FIXME: should cache?
  var ret MethodTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&trans), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MethodTweener) SetEase(ease TweenEaseType, ) MethodTweener {
  classNameV := StringNameFromStr("MethodTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ease")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 315540545) // FIXME: should cache?
  var ret MethodTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ease), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals

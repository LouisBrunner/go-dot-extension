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

type MethodTweener struct {
  Tweener
}

func (me *MethodTweener) BaseClass() string {
  return "MethodTweener"
}

func NewMethodTweener() *MethodTweener {
  str := StringNameFromStr("MethodTweener") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MethodTweener{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *MethodTweener) SetDelay(delay float64, ) MethodTweener {
  classNameV := StringNameFromStr("MethodTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 266477812) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMethodTweener()
  pinner.Pin(&delay)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MethodTweener) SetTrans(trans TweenTransitionType, ) MethodTweener {
  classNameV := StringNameFromStr("MethodTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_trans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740975367) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&trans) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMethodTweener()
  pinner.Pin(&trans)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MethodTweener) SetEase(ease TweenEaseType, ) MethodTweener {
  classNameV := StringNameFromStr("MethodTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ease")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 315540545) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ease) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMethodTweener()
  pinner.Pin(&ease)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

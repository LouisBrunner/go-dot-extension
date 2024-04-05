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

type ptrsForMethodTweenerList struct {
  fnSetDelay gdc.MethodBindPtr
  fnSetTrans gdc.MethodBindPtr
  fnSetEase gdc.MethodBindPtr
}

var ptrsForMethodTweener ptrsForMethodTweenerList

func initMethodTweenerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("MethodTweener")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_delay")
    defer methodName.Destroy()
    ptrsForMethodTweener.fnSetDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 266477812))
  }
  {
    methodName := StringNameFromStr("set_trans")
    defer methodName.Destroy()
    ptrsForMethodTweener.fnSetTrans = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740975367))
  }
  {
    methodName := StringNameFromStr("set_ease")
    defer methodName.Destroy()
    ptrsForMethodTweener.fnSetEase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 315540545))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMethodTweener()
  pinner.Pin(&delay)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMethodTweener.fnSetDelay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MethodTweener) SetTrans(trans TweenTransitionType, ) MethodTweener {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&trans) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMethodTweener()
  pinner.Pin(&trans)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMethodTweener.fnSetTrans), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MethodTweener) SetEase(ease TweenEaseType, ) MethodTweener {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ease) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMethodTweener()
  pinner.Pin(&ease)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMethodTweener.fnSetEase), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

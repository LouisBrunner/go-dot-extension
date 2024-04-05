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

type ptrsForPropertyTweenerList struct {
  fnFrom gdc.MethodBindPtr
  fnFromCurrent gdc.MethodBindPtr
  fnAsRelative gdc.MethodBindPtr
  fnSetTrans gdc.MethodBindPtr
  fnSetEase gdc.MethodBindPtr
  fnSetDelay gdc.MethodBindPtr
}

var ptrsForPropertyTweener ptrsForPropertyTweenerList

func initPropertyTweenerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PropertyTweener")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("from")
    defer methodName.Destroy()
    ptrsForPropertyTweener.fnFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4190193059))
  }
  {
    methodName := StringNameFromStr("from_current")
    defer methodName.Destroy()
    ptrsForPropertyTweener.fnFromCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4279177709))
  }
  {
    methodName := StringNameFromStr("as_relative")
    defer methodName.Destroy()
    ptrsForPropertyTweener.fnAsRelative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4279177709))
  }
  {
    methodName := StringNameFromStr("set_trans")
    defer methodName.Destroy()
    ptrsForPropertyTweener.fnSetTrans = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1899107404))
  }
  {
    methodName := StringNameFromStr("set_ease")
    defer methodName.Destroy()
    ptrsForPropertyTweener.fnSetEase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1080455622))
  }
  {
    methodName := StringNameFromStr("set_delay")
    defer methodName.Destroy()
    ptrsForPropertyTweener.fnSetDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2171559331))
  }
}

type PropertyTweener struct {
  Tweener
}

func (me *PropertyTweener) BaseClass() string {
  return "PropertyTweener"
}

func NewPropertyTweener() *PropertyTweener {
  str := StringNameFromStr("PropertyTweener") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PropertyTweener{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PropertyTweener) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PropertyTweener) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PropertyTweener) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PropertyTweener) From(value Variant, ) PropertyTweener {
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPropertyTweener()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPropertyTweener.fnFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) FromCurrent() PropertyTweener {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPropertyTweener()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPropertyTweener.fnFromCurrent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) AsRelative() PropertyTweener {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPropertyTweener()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPropertyTweener.fnAsRelative), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) SetTrans(trans TweenTransitionType, ) PropertyTweener {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&trans) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPropertyTweener()
  pinner.Pin(&trans)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPropertyTweener.fnSetTrans), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) SetEase(ease TweenEaseType, ) PropertyTweener {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ease) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPropertyTweener()
  pinner.Pin(&ease)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPropertyTweener.fnSetEase), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) SetDelay(delay float64, ) PropertyTweener {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPropertyTweener()
  pinner.Pin(&delay)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPropertyTweener.fnSetDelay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

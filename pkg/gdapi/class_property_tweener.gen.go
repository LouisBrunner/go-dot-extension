// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4190193059) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(value.AsCTypePtr()), }
  ret := NewPropertyTweener()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) FromCurrent() PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4279177709) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPropertyTweener()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) AsRelative() PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("as_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4279177709) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPropertyTweener()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) SetTrans(trans TweenTransitionType, ) PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_trans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1899107404) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&trans), }
  ret := NewPropertyTweener()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) SetEase(ease TweenEaseType, ) PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ease")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1080455622) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ease), }
  ret := NewPropertyTweener()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PropertyTweener) SetDelay(delay float64, ) PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2171559331) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay), }
  ret := NewPropertyTweener()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

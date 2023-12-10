// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PropertyTweener struct {
  obj gdc.ObjectPtr
}

func (me *PropertyTweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PropertyTweener) BaseClass() string {
  return "PropertyTweener"
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
  var ret PropertyTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PropertyTweener) FromCurrent() PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4279177709) // FIXME: should cache?
  var ret PropertyTweener
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PropertyTweener) AsRelative() PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("as_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4279177709) // FIXME: should cache?
  var ret PropertyTweener
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PropertyTweener) SetTrans(trans TweenTransitionType, ) PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_trans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1899107404) // FIXME: should cache?
  var ret PropertyTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&trans), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PropertyTweener) SetEase(ease TweenEaseType, ) PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ease")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1080455622) // FIXME: should cache?
  var ret PropertyTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ease), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PropertyTweener) SetDelay(delay float32, ) PropertyTweener {
  classNameV := StringNameFromStr("PropertyTweener")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2171559331) // FIXME: should cache?
  var ret PropertyTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties
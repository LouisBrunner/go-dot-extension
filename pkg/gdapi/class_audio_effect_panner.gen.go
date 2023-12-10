// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectPanner struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectPanner) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectPanner) BaseClass() string {
  return "AudioEffectPanner"
}



// Enums

func (me *AudioEffectPanner) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectPanner) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectPanner) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectPanner) SetPan(cpanume float32, )  {
  classNameV := StringNameFromStr("AudioEffectPanner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cpanume), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectPanner) GetPan() float32 {
  classNameV := StringNameFromStr("AudioEffectPanner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AudioEffectPanner) GetPropPan() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectPanner) SetPropPan(value float32) {
  panic("TODO: implement")
}
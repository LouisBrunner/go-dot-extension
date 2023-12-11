// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioStreamPolyphonic struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamPolyphonic) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamPolyphonic) BaseClass() string {
  return "AudioStreamPolyphonic"
}



// Enums

func (me *AudioStreamPolyphonic) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPolyphonic) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPolyphonic) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamPolyphonic) SetPolyphony(voices int, )  {
  classNameV := StringNameFromStr("AudioStreamPolyphonic")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voices), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPolyphonic) GetPolyphony() int {
  classNameV := StringNameFromStr("AudioStreamPolyphonic")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

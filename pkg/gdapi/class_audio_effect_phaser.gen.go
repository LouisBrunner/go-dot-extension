// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectPhaser struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectPhaser) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectPhaser) BaseClass() string {
  return "AudioEffectPhaser"
}



// Enums

func (me *AudioEffectPhaser) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectPhaser) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectPhaser) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectPhaser) SetRangeMinHz(hz float32, )  {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_range_min_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectPhaser) GetRangeMinHz() float32 {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_range_min_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectPhaser) SetRangeMaxHz(hz float32, )  {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_range_max_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectPhaser) GetRangeMaxHz() float32 {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_range_max_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectPhaser) SetRateHz(hz float32, )  {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rate_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectPhaser) GetRateHz() float32 {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rate_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectPhaser) SetFeedback(fbk float32, )  {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fbk), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectPhaser) GetFeedback() float32 {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectPhaser) SetDepth(depth float32, )  {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectPhaser) GetDepth() float32 {
  classNameV := StringNameFromStr("AudioEffectPhaser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

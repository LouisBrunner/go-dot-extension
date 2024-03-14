// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectEQ struct {
  AudioEffect
}

func (me *AudioEffectEQ) BaseClass() string {
  return "AudioEffectEQ"
}



// Enums

func (me *AudioEffectEQ) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectEQ) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectEQ) SetBandGainDb(band_idx int, volume_db float32, )  {
  classNameV := StringNameFromStr("AudioEffectEQ")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_band_gain_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&band_idx), gdc.ConstTypePtr(&volume_db), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectEQ) GetBandGainDb(band_idx int, ) float32 {
  classNameV := StringNameFromStr("AudioEffectEQ")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_band_gain_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&band_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectEQ) GetBandCount() int {
  classNameV := StringNameFromStr("AudioEffectEQ")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_band_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals

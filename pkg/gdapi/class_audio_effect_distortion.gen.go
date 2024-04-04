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

type AudioEffectDistortion struct {
  AudioEffect
}

func (me *AudioEffectDistortion) BaseClass() string {
  return "AudioEffectDistortion"
}

func NewAudioEffectDistortion() *AudioEffectDistortion {
  str := StringNameFromStr("AudioEffectDistortion") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectDistortion{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AudioEffectDistortionMode int
const (
  AudioEffectDistortionModeModeClip AudioEffectDistortionMode = 0
  AudioEffectDistortionModeModeAtan AudioEffectDistortionMode = 1
  AudioEffectDistortionModeModeLofi AudioEffectDistortionMode = 2
  AudioEffectDistortionModeModeOverdrive AudioEffectDistortionMode = 3
  AudioEffectDistortionModeModeWaveshape AudioEffectDistortionMode = 4
)

func (me *AudioEffectDistortion) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectDistortion) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectDistortion) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectDistortion) SetMode(mode AudioEffectDistortionMode, )  {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1314744793) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetMode() AudioEffectDistortionMode {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 809118343) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioEffectDistortionMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AudioEffectDistortion) SetPreGain(pre_gain float64, )  {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pre_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pre_gain) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetPreGain() float64 {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pre_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectDistortion) SetKeepHfHz(keep_hf_hz float64, )  {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_hf_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep_hf_hz) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetKeepHfHz() float64 {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keep_hf_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectDistortion) SetDrive(drive float64, )  {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&drive) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetDrive() float64 {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectDistortion) SetPostGain(post_gain float64, )  {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_post_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&post_gain) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetPostGain() float64 {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_post_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

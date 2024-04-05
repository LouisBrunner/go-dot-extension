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

type ptrsForAudioEffectDistortionList struct {
  fnSetMode gdc.MethodBindPtr
  fnGetMode gdc.MethodBindPtr
  fnSetPreGain gdc.MethodBindPtr
  fnGetPreGain gdc.MethodBindPtr
  fnSetKeepHfHz gdc.MethodBindPtr
  fnGetKeepHfHz gdc.MethodBindPtr
  fnSetDrive gdc.MethodBindPtr
  fnGetDrive gdc.MethodBindPtr
  fnSetPostGain gdc.MethodBindPtr
  fnGetPostGain gdc.MethodBindPtr
}

var ptrsForAudioEffectDistortion ptrsForAudioEffectDistortionList

func initAudioEffectDistortionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioEffectDistortion")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_mode")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1314744793))
  }
  {
    methodName := StringNameFromStr("get_mode")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 809118343))
  }
  {
    methodName := StringNameFromStr("set_pre_gain")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnSetPreGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_pre_gain")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnGetPreGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_keep_hf_hz")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnSetKeepHfHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_keep_hf_hz")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnGetKeepHfHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_drive")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnSetDrive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_drive")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnGetDrive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_post_gain")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnSetPostGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_post_gain")
    defer methodName.Destroy()
    ptrsForAudioEffectDistortion.fnGetPostGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetMode() AudioEffectDistortionMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioEffectDistortionMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AudioEffectDistortion) SetPreGain(pre_gain float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pre_gain) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnSetPreGain), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetPreGain() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnGetPreGain), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectDistortion) SetKeepHfHz(keep_hf_hz float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep_hf_hz) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnSetKeepHfHz), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetKeepHfHz() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnGetKeepHfHz), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectDistortion) SetDrive(drive float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&drive) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnSetDrive), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetDrive() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnGetDrive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectDistortion) SetPostGain(post_gain float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&post_gain) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnSetPostGain), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectDistortion) GetPostGain() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDistortion.fnGetPostGain), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

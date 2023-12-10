// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectDistortion struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectDistortion) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectDistortion) BaseClass() string {
  return "AudioEffectDistortion"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDistortion) GetMode() AudioEffectDistortionMode {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 809118343) // FIXME: should cache?
  var ret AudioEffectDistortionMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDistortion) SetPreGain(pre_gain float32, )  {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pre_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pre_gain), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDistortion) GetPreGain() float32 {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pre_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDistortion) SetKeepHfHz(keep_hf_hz float32, )  {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_hf_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep_hf_hz), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDistortion) GetKeepHfHz() float32 {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keep_hf_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDistortion) SetDrive(drive float32, )  {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&drive), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDistortion) GetDrive() float32 {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDistortion) SetPostGain(post_gain float32, )  {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_post_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&post_gain), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDistortion) GetPostGain() float32 {
  classNameV := StringNameFromStr("AudioEffectDistortion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_post_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AudioEffectDistortion) GetPropMode() int {
  panic("TODO: implement")
}

func (me *AudioEffectDistortion) SetPropMode(value int) {
  panic("TODO: implement")
}

func (me *AudioEffectDistortion) GetPropPreGain() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectDistortion) SetPropPreGain(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectDistortion) GetPropKeepHfHz() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectDistortion) SetPropKeepHfHz(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectDistortion) GetPropDrive() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectDistortion) SetPropDrive(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectDistortion) GetPropPostGain() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectDistortion) SetPropPostGain(value float32) {
  panic("TODO: implement")
}
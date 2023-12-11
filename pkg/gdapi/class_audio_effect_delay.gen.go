// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectDelay struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectDelay) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectDelay) BaseClass() string {
  return "AudioEffectDelay"
}



// Enums

func (me *AudioEffectDelay) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectDelay) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectDelay) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectDelay) SetDry(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dry")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetDry() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dry")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetTap1Active(amount bool, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap1_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) IsTap1Active() bool {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_tap1_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetTap1DelayMs(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap1_delay_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetTap1DelayMs() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tap1_delay_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetTap1LevelDb(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap1_level_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetTap1LevelDb() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tap1_level_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetTap1Pan(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap1_pan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetTap1Pan() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tap1_pan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetTap2Active(amount bool, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap2_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) IsTap2Active() bool {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_tap2_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetTap2DelayMs(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap2_delay_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetTap2DelayMs() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tap2_delay_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetTap2LevelDb(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap2_level_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetTap2LevelDb() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tap2_level_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetTap2Pan(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap2_pan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetTap2Pan() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tap2_pan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetFeedbackActive(amount bool, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_feedback_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) IsFeedbackActive() bool {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_feedback_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetFeedbackDelayMs(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_feedback_delay_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetFeedbackDelayMs() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feedback_delay_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetFeedbackLevelDb(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_feedback_level_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetFeedbackLevelDb() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feedback_level_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectDelay) SetFeedbackLowpass(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_feedback_lowpass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectDelay) GetFeedbackLowpass() float32 {
  classNameV := StringNameFromStr("AudioEffectDelay")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feedback_lowpass")
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

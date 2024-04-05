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

type ptrsForAudioEffectReverbList struct {
  fnSetPredelayMsec gdc.MethodBindPtr
  fnGetPredelayMsec gdc.MethodBindPtr
  fnSetPredelayFeedback gdc.MethodBindPtr
  fnGetPredelayFeedback gdc.MethodBindPtr
  fnSetRoomSize gdc.MethodBindPtr
  fnGetRoomSize gdc.MethodBindPtr
  fnSetDamping gdc.MethodBindPtr
  fnGetDamping gdc.MethodBindPtr
  fnSetSpread gdc.MethodBindPtr
  fnGetSpread gdc.MethodBindPtr
  fnSetDry gdc.MethodBindPtr
  fnGetDry gdc.MethodBindPtr
  fnSetWet gdc.MethodBindPtr
  fnGetWet gdc.MethodBindPtr
  fnSetHpf gdc.MethodBindPtr
  fnGetHpf gdc.MethodBindPtr
}

var ptrsForAudioEffectReverb ptrsForAudioEffectReverbList

func initAudioEffectReverbPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioEffectReverb")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_predelay_msec")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnSetPredelayMsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_predelay_msec")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnGetPredelayMsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_predelay_feedback")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnSetPredelayFeedback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_predelay_feedback")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnGetPredelayFeedback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_room_size")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnSetRoomSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_room_size")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnGetRoomSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_damping")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnSetDamping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_damping")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnGetDamping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_spread")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnSetSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_spread")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnGetSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_dry")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnSetDry = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_dry")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnGetDry = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_wet")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnSetWet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_wet")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnGetWet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_hpf")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnSetHpf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_hpf")
    defer methodName.Destroy()
    ptrsForAudioEffectReverb.fnGetHpf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type AudioEffectReverb struct {
  AudioEffect
}

func (me *AudioEffectReverb) BaseClass() string {
  return "AudioEffectReverb"
}

func NewAudioEffectReverb() *AudioEffectReverb {
  str := StringNameFromStr("AudioEffectReverb") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectReverb{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectReverb) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectReverb) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectReverb) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectReverb) SetPredelayMsec(msec float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnSetPredelayMsec), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectReverb) GetPredelayMsec() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnGetPredelayMsec), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectReverb) SetPredelayFeedback(feedback float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feedback) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnSetPredelayFeedback), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectReverb) GetPredelayFeedback() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnGetPredelayFeedback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectReverb) SetRoomSize(size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnSetRoomSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectReverb) GetRoomSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnGetRoomSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectReverb) SetDamping(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnSetDamping), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectReverb) GetDamping() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnGetDamping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectReverb) SetSpread(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnSetSpread), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectReverb) GetSpread() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnGetSpread), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectReverb) SetDry(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnSetDry), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectReverb) GetDry() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnGetDry), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectReverb) SetWet(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnSetWet), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectReverb) GetWet() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnGetWet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectReverb) SetHpf(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnSetHpf), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectReverb) GetHpf() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectReverb.fnGetHpf), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type AudioEffectCapture struct {
  AudioEffect
}

func (me *AudioEffectCapture) BaseClass() string {
  return "AudioEffectCapture"
}

func NewAudioEffectCapture() *AudioEffectCapture {
  str := StringNameFromStr("AudioEffectCapture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectCapture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectCapture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectCapture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectCapture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectCapture) CanGetBuffer(frames int64, ) bool {
  classNameV := StringNameFromStr("AudioEffectCapture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_get_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&frames)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectCapture) GetBuffer(frames int64, ) PackedVector2Array {
  classNameV := StringNameFromStr("AudioEffectCapture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2649534757) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()
  pinner.Pin(&frames)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioEffectCapture) ClearBuffer()  {
  classNameV := StringNameFromStr("AudioEffectCapture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectCapture) SetBufferLength(buffer_length_seconds float64, )  {
  classNameV := StringNameFromStr("AudioEffectCapture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffer_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_length_seconds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectCapture) GetBufferLength() float64 {
  classNameV := StringNameFromStr("AudioEffectCapture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectCapture) GetFramesAvailable() int64 {
  classNameV := StringNameFromStr("AudioEffectCapture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frames_available")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectCapture) GetDiscardedFrames() int64 {
  classNameV := StringNameFromStr("AudioEffectCapture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_discarded_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectCapture) GetBufferLengthFrames() int64 {
  classNameV := StringNameFromStr("AudioEffectCapture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer_length_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectCapture) GetPushedFrames() int64 {
  classNameV := StringNameFromStr("AudioEffectCapture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pushed_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioStreamGeneratorPlayback struct {
  AudioStreamPlaybackResampled
}

func (me *AudioStreamGeneratorPlayback) BaseClass() string {
  return "AudioStreamGeneratorPlayback"
}

func NewAudioStreamGeneratorPlayback() *AudioStreamGeneratorPlayback {
  str := StringNameFromStr("AudioStreamGeneratorPlayback") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamGeneratorPlayback{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioStreamGeneratorPlayback) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamGeneratorPlayback) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamGeneratorPlayback) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamGeneratorPlayback) PushFrame(frame Vector2, ) bool {
  classNameV := StringNameFromStr("AudioStreamGeneratorPlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3975407249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(frame.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) CanPushBuffer(amount int64, ) bool {
  classNameV := StringNameFromStr("AudioStreamGeneratorPlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_push_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) PushBuffer(frames PackedVector2Array, ) bool {
  classNameV := StringNameFromStr("AudioStreamGeneratorPlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1361156557) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(frames.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) GetFramesAvailable() int64 {
  classNameV := StringNameFromStr("AudioStreamGeneratorPlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frames_available")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) GetSkips() int64 {
  classNameV := StringNameFromStr("AudioStreamGeneratorPlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skips")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) ClearBuffer()  {
  classNameV := StringNameFromStr("AudioStreamGeneratorPlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

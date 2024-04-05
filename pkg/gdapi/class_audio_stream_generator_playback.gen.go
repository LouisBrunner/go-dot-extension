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

type ptrsForAudioStreamGeneratorPlaybackList struct {
  fnPushFrame gdc.MethodBindPtr
  fnCanPushBuffer gdc.MethodBindPtr
  fnPushBuffer gdc.MethodBindPtr
  fnGetFramesAvailable gdc.MethodBindPtr
  fnGetSkips gdc.MethodBindPtr
  fnClearBuffer gdc.MethodBindPtr
}

var ptrsForAudioStreamGeneratorPlayback ptrsForAudioStreamGeneratorPlaybackList

func initAudioStreamGeneratorPlaybackPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioStreamGeneratorPlayback")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("push_frame")
    defer methodName.Destroy()
    ptrsForAudioStreamGeneratorPlayback.fnPushFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3975407249))
  }
  {
    methodName := StringNameFromStr("can_push_buffer")
    defer methodName.Destroy()
    ptrsForAudioStreamGeneratorPlayback.fnCanPushBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("push_buffer")
    defer methodName.Destroy()
    ptrsForAudioStreamGeneratorPlayback.fnPushBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1361156557))
  }
  {
    methodName := StringNameFromStr("get_frames_available")
    defer methodName.Destroy()
    ptrsForAudioStreamGeneratorPlayback.fnGetFramesAvailable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_skips")
    defer methodName.Destroy()
    ptrsForAudioStreamGeneratorPlayback.fnGetSkips = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("clear_buffer")
    defer methodName.Destroy()
    ptrsForAudioStreamGeneratorPlayback.fnClearBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

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
  cargs := []gdc.ConstTypePtr{frame.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGeneratorPlayback.fnPushFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) CanPushBuffer(amount int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&amount)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGeneratorPlayback.fnCanPushBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) PushBuffer(frames PackedVector2Array, ) bool {
  cargs := []gdc.ConstTypePtr{frames.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGeneratorPlayback.fnPushBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) GetFramesAvailable() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGeneratorPlayback.fnGetFramesAvailable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) GetSkips() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGeneratorPlayback.fnGetSkips), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGeneratorPlayback) ClearBuffer()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGeneratorPlayback.fnClearBuffer), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

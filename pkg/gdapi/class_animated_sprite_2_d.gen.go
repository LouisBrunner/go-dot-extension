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

type ptrsForAnimatedSprite2DList struct {
  fnSetSpriteFrames gdc.MethodBindPtr
  fnGetSpriteFrames gdc.MethodBindPtr
  fnSetAnimation gdc.MethodBindPtr
  fnGetAnimation gdc.MethodBindPtr
  fnSetAutoplay gdc.MethodBindPtr
  fnGetAutoplay gdc.MethodBindPtr
  fnIsPlaying gdc.MethodBindPtr
  fnPlay gdc.MethodBindPtr
  fnPlayBackwards gdc.MethodBindPtr
  fnPause gdc.MethodBindPtr
  fnStop gdc.MethodBindPtr
  fnSetCentered gdc.MethodBindPtr
  fnIsCentered gdc.MethodBindPtr
  fnSetOffset gdc.MethodBindPtr
  fnGetOffset gdc.MethodBindPtr
  fnSetFlipH gdc.MethodBindPtr
  fnIsFlippedH gdc.MethodBindPtr
  fnSetFlipV gdc.MethodBindPtr
  fnIsFlippedV gdc.MethodBindPtr
  fnSetFrame gdc.MethodBindPtr
  fnGetFrame gdc.MethodBindPtr
  fnSetFrameProgress gdc.MethodBindPtr
  fnGetFrameProgress gdc.MethodBindPtr
  fnSetFrameAndProgress gdc.MethodBindPtr
  fnSetSpeedScale gdc.MethodBindPtr
  fnGetSpeedScale gdc.MethodBindPtr
  fnGetPlayingSpeed gdc.MethodBindPtr
}

var ptrsForAnimatedSprite2D ptrsForAnimatedSprite2DList

func initAnimatedSprite2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AnimatedSprite2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_sprite_frames")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetSpriteFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 905781144))
  }
  {
    methodName := StringNameFromStr("get_sprite_frames")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnGetSpriteFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3804851214))
  }
  {
    methodName := StringNameFromStr("set_animation")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_animation")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnGetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_autoplay")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_autoplay")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnGetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("is_playing")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnIsPlaying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("play")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnPlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2372066587))
  }
  {
    methodName := StringNameFromStr("play_backwards")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnPlayBackwards = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1421762485))
  }
  {
    methodName := StringNameFromStr("pause")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnPause = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("stop")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_centered")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_centered")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnIsCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_offset")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_offset")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_flip_h")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetFlipH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_flipped_h")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnIsFlippedH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_flip_v")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetFlipV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_flipped_v")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnIsFlippedV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_frame")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_frame")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnGetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_frame_progress")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetFrameProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_frame_progress")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnGetFrameProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_frame_and_progress")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetFrameAndProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("set_speed_scale")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_speed_scale")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnGetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_playing_speed")
    defer methodName.Destroy()
    ptrsForAnimatedSprite2D.fnGetPlayingSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type AnimatedSprite2D struct {
  Node2D
}

func (me *AnimatedSprite2D) BaseClass() string {
  return "AnimatedSprite2D"
}

func NewAnimatedSprite2D() *AnimatedSprite2D {
  str := StringNameFromStr("AnimatedSprite2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimatedSprite2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AnimatedSprite2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimatedSprite2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimatedSprite2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimatedSprite2D) SetSpriteFrames(sprite_frames SpriteFrames, )  {
  cargs := []gdc.ConstTypePtr{sprite_frames.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetSpriteFrames), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) GetSpriteFrames() SpriteFrames {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSpriteFrames()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnGetSpriteFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimatedSprite2D) SetAnimation(name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) GetAnimation() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnGetAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimatedSprite2D) SetAutoplay(name String, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetAutoplay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) GetAutoplay() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnGetAutoplay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimatedSprite2D) IsPlaying() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnIsPlaying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedSprite2D) Play(name StringName, custom_speed float64, from_end bool, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&custom_speed) , gdc.ConstTypePtr(&from_end) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnPlay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) PlayBackwards(name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnPlayBackwards), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) Pause()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnPause), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) Stop()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) SetCentered(centered bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&centered) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetCentered), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) IsCentered() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnIsCentered), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedSprite2D) SetOffset(offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) GetOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimatedSprite2D) SetFlipH(flip_h bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_h) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetFlipH), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) IsFlippedH() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnIsFlippedH), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedSprite2D) SetFlipV(flip_v bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_v) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetFlipV), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) IsFlippedV() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnIsFlippedV), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedSprite2D) SetFrame(frame int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetFrame), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) GetFrame() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnGetFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedSprite2D) SetFrameProgress(progress float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&progress) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetFrameProgress), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) GetFrameProgress() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnGetFrameProgress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedSprite2D) SetFrameAndProgress(frame int64, progress float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame) , gdc.ConstTypePtr(&progress) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetFrameAndProgress), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) SetSpeedScale(speed_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnSetSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedSprite2D) GetSpeedScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnGetSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedSprite2D) GetPlayingSpeed() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite2D.fnGetPlayingSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimatedSprite2DSpriteFramesChangedSignalFn func()

func (me *AnimatedSprite2D) ConnectSpriteFramesChanged(subs SignalSubscribers, fn AnimatedSprite2DSpriteFramesChangedSignalFn) {
  sig := StringNameFromStr("sprite_frames_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite2D) DisconnectSpriteFramesChanged(subs SignalSubscribers, fn AnimatedSprite2DSpriteFramesChangedSignalFn) {
  sig := StringNameFromStr("sprite_frames_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimatedSprite2DAnimationChangedSignalFn func()

func (me *AnimatedSprite2D) ConnectAnimationChanged(subs SignalSubscribers, fn AnimatedSprite2DAnimationChangedSignalFn) {
  sig := StringNameFromStr("animation_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite2D) DisconnectAnimationChanged(subs SignalSubscribers, fn AnimatedSprite2DAnimationChangedSignalFn) {
  sig := StringNameFromStr("animation_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimatedSprite2DFrameChangedSignalFn func()

func (me *AnimatedSprite2D) ConnectFrameChanged(subs SignalSubscribers, fn AnimatedSprite2DFrameChangedSignalFn) {
  sig := StringNameFromStr("frame_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite2D) DisconnectFrameChanged(subs SignalSubscribers, fn AnimatedSprite2DFrameChangedSignalFn) {
  sig := StringNameFromStr("frame_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimatedSprite2DAnimationLoopedSignalFn func()

func (me *AnimatedSprite2D) ConnectAnimationLooped(subs SignalSubscribers, fn AnimatedSprite2DAnimationLoopedSignalFn) {
  sig := StringNameFromStr("animation_looped")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite2D) DisconnectAnimationLooped(subs SignalSubscribers, fn AnimatedSprite2DAnimationLoopedSignalFn) {
  sig := StringNameFromStr("animation_looped")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimatedSprite2DAnimationFinishedSignalFn func()

func (me *AnimatedSprite2D) ConnectAnimationFinished(subs SignalSubscribers, fn AnimatedSprite2DAnimationFinishedSignalFn) {
  sig := StringNameFromStr("animation_finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite2D) DisconnectAnimationFinished(subs SignalSubscribers, fn AnimatedSprite2DAnimationFinishedSignalFn) {
  sig := StringNameFromStr("animation_finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

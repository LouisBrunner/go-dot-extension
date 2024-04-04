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

type AnimationPlayer struct {
  AnimationMixer
}

func (me *AnimationPlayer) BaseClass() string {
  return "AnimationPlayer"
}

func NewAnimationPlayer() *AnimationPlayer {
  str := StringNameFromStr("AnimationPlayer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationPlayer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AnimationPlayerAnimationProcessCallback int
const (
  AnimationPlayerAnimationProcessCallbackAnimationProcessPhysics AnimationPlayerAnimationProcessCallback = 0
  AnimationPlayerAnimationProcessCallbackAnimationProcessIdle AnimationPlayerAnimationProcessCallback = 1
  AnimationPlayerAnimationProcessCallbackAnimationProcessManual AnimationPlayerAnimationProcessCallback = 2
)

type AnimationPlayerAnimationMethodCallMode int
const (
  AnimationPlayerAnimationMethodCallModeAnimationMethodCallDeferred AnimationPlayerAnimationMethodCallMode = 0
  AnimationPlayerAnimationMethodCallModeAnimationMethodCallImmediate AnimationPlayerAnimationMethodCallMode = 1
)

func (me *AnimationPlayer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationPlayer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationPlayer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationPlayer) AnimationSetNext(animation_from StringName, animation_to StringName, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("animation_set_next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{animation_from.AsCTypePtr(), animation_to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) AnimationGetNext(animation_from StringName, ) StringName {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("animation_get_next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965194235) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{animation_from.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationPlayer) SetBlendTime(animation_from StringName, animation_to StringName, sec float64, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3231131886) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{animation_from.AsCTypePtr(), animation_to.AsCTypePtr(), gdc.ConstTypePtr(&sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetBlendTime(animation_from StringName, animation_to StringName, ) float64 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1958752504) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{animation_from.AsCTypePtr(), animation_to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationPlayer) SetDefaultBlendTime(sec float64, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_blend_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetDefaultBlendTime() float64 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_blend_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationPlayer) Play(name StringName, custom_blend float64, custom_speed float64, from_end bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3118260607) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&custom_blend) , gdc.ConstTypePtr(&custom_speed) , gdc.ConstTypePtr(&from_end) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) PlayBackwards(name StringName, custom_blend float64, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_backwards")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2787282401) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&custom_blend) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) Pause()  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pause")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) Stop(keep_state bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 107499316) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep_state) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) IsPlaying() bool {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationPlayer) SetCurrentAnimation(animation String, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{animation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetCurrentAnimation() String {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationPlayer) SetAssignedAnimation(animation String, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_assigned_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{animation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetAssignedAnimation() String {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_assigned_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationPlayer) Queue(name StringName, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetQueue() PackedStringArray {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_queue")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationPlayer) ClearQueue()  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_queue")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) SetSpeedScale(speed float64, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetSpeedScale() float64 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationPlayer) GetPlayingSpeed() float64 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playing_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationPlayer) SetAutoplay(name String, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetAutoplay() String {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationPlayer) SetMovieQuitOnFinishEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_movie_quit_on_finish_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) IsMovieQuitOnFinishEnabled() bool {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_movie_quit_on_finish_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationPlayer) GetCurrentAnimationPosition() float64 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_animation_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationPlayer) GetCurrentAnimationLength() float64 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_animation_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationPlayer) Seek(seconds float64, update bool, update_only bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("seek")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1807872683) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds) , gdc.ConstTypePtr(&update) , gdc.ConstTypePtr(&update_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) SetProcessCallback(mode AnimationPlayerAnimationProcessCallback, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1663839457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetProcessCallback() AnimationPlayerAnimationProcessCallback {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4207496604) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationPlayerAnimationProcessCallback

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationPlayer) SetMethodCallMode(mode AnimationPlayerAnimationMethodCallMode, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_method_call_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3413514846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetMethodCallMode() AnimationPlayerAnimationMethodCallMode {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_method_call_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3583380054) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationPlayerAnimationMethodCallMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationPlayer) SetRoot(path NodePath, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationPlayer) GetRoot() NodePath {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationPlayerCurrentAnimationChangedSignalFn func(name String, )

func (me *AnimationPlayer) ConnectCurrentAnimationChanged(subs SignalSubscribers, fn AnimationPlayerCurrentAnimationChangedSignalFn) {
  sig := StringNameFromStr("current_animation_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectCurrentAnimationChanged(subs SignalSubscribers, fn AnimationPlayerCurrentAnimationChangedSignalFn) {
  sig := StringNameFromStr("current_animation_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationPlayerAnimationChangedSignalFn func(old_name StringName, new_name StringName, )

func (me *AnimationPlayer) ConnectAnimationChanged(subs SignalSubscribers, fn AnimationPlayerAnimationChangedSignalFn) {
  sig := StringNameFromStr("animation_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectAnimationChanged(subs SignalSubscribers, fn AnimationPlayerAnimationChangedSignalFn) {
  sig := StringNameFromStr("animation_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

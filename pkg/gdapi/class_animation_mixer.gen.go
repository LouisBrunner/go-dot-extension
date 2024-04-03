// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationMixer struct {
  Node
}

func (me *AnimationMixer) BaseClass() string {
  return "AnimationMixer"
}

func NewAnimationMixer() *AnimationMixer {
  str := StringNameFromStr("AnimationMixer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationMixer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AnimationMixerAnimationCallbackModeProcess int
const (
  AnimationMixerAnimationCallbackModeProcessAnimationCallbackModeProcessPhysics AnimationMixerAnimationCallbackModeProcess = 0
  AnimationMixerAnimationCallbackModeProcessAnimationCallbackModeProcessIdle AnimationMixerAnimationCallbackModeProcess = 1
  AnimationMixerAnimationCallbackModeProcessAnimationCallbackModeProcessManual AnimationMixerAnimationCallbackModeProcess = 2
)

type AnimationMixerAnimationCallbackModeMethod int
const (
  AnimationMixerAnimationCallbackModeMethodAnimationCallbackModeMethodDeferred AnimationMixerAnimationCallbackModeMethod = 0
  AnimationMixerAnimationCallbackModeMethodAnimationCallbackModeMethodImmediate AnimationMixerAnimationCallbackModeMethod = 1
)

func (me *AnimationMixer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationMixer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationMixer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationMixer) AddAnimationLibrary(name StringName, library AnimationLibrary, ) Error {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 618909818) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(library.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationMixer) RemoveAnimationLibrary(name StringName, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) RenameAnimationLibrary(name StringName, newname StringName, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(newname.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) HasAnimationLibrary(name StringName, ) bool {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationMixer) GetAnimationLibrary(name StringName, ) AnimationLibrary {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 147342321) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  ret := NewAnimationLibrary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) GetAnimationLibraryList() []StringName {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_library_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[StringName](ret)
}

func  (me *AnimationMixer) HasAnimation(name StringName, ) bool {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationMixer) GetAnimation(name StringName, ) Animation {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2933122410) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  ret := NewAnimation()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) GetAnimationList() PackedStringArray {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) SetActive(active bool, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) IsActive() bool {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationMixer) SetDeterministic(deterministic bool, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deterministic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&deterministic), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) IsDeterministic() bool {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_deterministic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationMixer) SetRootNode(path NodePath, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) GetRootNode() NodePath {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) SetCallbackModeProcess(mode AnimationMixerAnimationCallbackModeProcess, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_callback_mode_process")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2153733086) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) GetCallbackModeProcess() AnimationMixerAnimationCallbackModeProcess {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_callback_mode_process")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1394468472) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret AnimationMixerAnimationCallbackModeProcess

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationMixer) SetCallbackModeMethod(mode AnimationMixerAnimationCallbackModeMethod, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_callback_mode_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 742218271) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) GetCallbackModeMethod() AnimationMixerAnimationCallbackModeMethod {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_callback_mode_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 489449656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret AnimationMixerAnimationCallbackModeMethod

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationMixer) SetAudioMaxPolyphony(max_polyphony int64, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_audio_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_polyphony), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) GetAudioMaxPolyphony() int64 {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_audio_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationMixer) SetRootMotionTrack(path NodePath, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_motion_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) GetRootMotionTrack() NodePath {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) GetRootMotionPosition() Vector3 {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) GetRootMotionRotation() Quaternion {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1222331677) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewQuaternion()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) GetRootMotionScale() Vector3 {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) GetRootMotionPositionAccumulator() Vector3 {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_position_accumulator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) GetRootMotionRotationAccumulator() Quaternion {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_rotation_accumulator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1222331677) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewQuaternion()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) GetRootMotionScaleAccumulator() Vector3 {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_scale_accumulator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) ClearCaches()  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_caches")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) Advance(delta float64, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("advance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) SetResetOnSaveEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reset_on_save_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationMixer) IsResetOnSaveEnabled() bool {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_reset_on_save_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationMixer) FindAnimation(animation Animation, ) StringName {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1559484580) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(animation.AsCTypePtr()), }
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationMixer) FindAnimationLibrary(animation Animation, ) StringName {
  classNameV := StringNameFromStr("AnimationMixer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1559484580) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(animation.AsCTypePtr()), }
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationMixerMixerUpdatedSignalFn func()

func (me *AnimationMixer) ConnectMixerUpdated(subs SignalSubscribers, fn AnimationMixerMixerUpdatedSignalFn) {
  sig := StringNameFromStr("mixer_updated")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectMixerUpdated(subs SignalSubscribers, fn AnimationMixerMixerUpdatedSignalFn) {
  sig := StringNameFromStr("mixer_updated")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerAnimationListChangedSignalFn func()

func (me *AnimationMixer) ConnectAnimationListChanged(subs SignalSubscribers, fn AnimationMixerAnimationListChangedSignalFn) {
  sig := StringNameFromStr("animation_list_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectAnimationListChanged(subs SignalSubscribers, fn AnimationMixerAnimationListChangedSignalFn) {
  sig := StringNameFromStr("animation_list_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerAnimationLibrariesUpdatedSignalFn func()

func (me *AnimationMixer) ConnectAnimationLibrariesUpdated(subs SignalSubscribers, fn AnimationMixerAnimationLibrariesUpdatedSignalFn) {
  sig := StringNameFromStr("animation_libraries_updated")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectAnimationLibrariesUpdated(subs SignalSubscribers, fn AnimationMixerAnimationLibrariesUpdatedSignalFn) {
  sig := StringNameFromStr("animation_libraries_updated")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerAnimationFinishedSignalFn func(anim_name StringName, )

func (me *AnimationMixer) ConnectAnimationFinished(subs SignalSubscribers, fn AnimationMixerAnimationFinishedSignalFn) {
  sig := StringNameFromStr("animation_finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectAnimationFinished(subs SignalSubscribers, fn AnimationMixerAnimationFinishedSignalFn) {
  sig := StringNameFromStr("animation_finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerAnimationStartedSignalFn func(anim_name StringName, )

func (me *AnimationMixer) ConnectAnimationStarted(subs SignalSubscribers, fn AnimationMixerAnimationStartedSignalFn) {
  sig := StringNameFromStr("animation_started")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectAnimationStarted(subs SignalSubscribers, fn AnimationMixerAnimationStartedSignalFn) {
  sig := StringNameFromStr("animation_started")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerCachesClearedSignalFn func()

func (me *AnimationMixer) ConnectCachesCleared(subs SignalSubscribers, fn AnimationMixerCachesClearedSignalFn) {
  sig := StringNameFromStr("caches_cleared")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectCachesCleared(subs SignalSubscribers, fn AnimationMixerCachesClearedSignalFn) {
  sig := StringNameFromStr("caches_cleared")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

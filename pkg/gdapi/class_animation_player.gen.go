// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationPlayer struct {
  obj gdc.ObjectPtr
}

func (me *AnimationPlayer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationPlayer) BaseClass() string {
  return "AnimationPlayer"
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

func  (me *AnimationPlayer) AddAnimationLibrary(name StringName, library AnimationLibrary, ) Error {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 618909818) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(library.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) RemoveAnimationLibrary(name StringName, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) RenameAnimationLibrary(name StringName, newname StringName, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(newname.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) HasAnimationLibrary(name StringName, ) bool {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) GetAnimationLibrary(name StringName, ) AnimationLibrary {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 147342321) // FIXME: should cache?
  var ret AnimationLibrary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) GetAnimationLibraryList() StringName {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_library_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) HasAnimation(name StringName, ) bool {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) GetAnimation(name StringName, ) Animation {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2933122410) // FIXME: should cache?
  var ret Animation
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) GetAnimationList() PackedStringArray {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) AnimationSetNext(anim_from StringName, anim_to StringName, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("animation_set_next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(anim_from.AsCTypePtr()), gdc.ConstTypePtr(anim_to.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) AnimationGetNext(anim_from StringName, ) StringName {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("animation_get_next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965194235) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(anim_from.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetBlendTime(anim_from StringName, anim_to StringName, sec float32, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3231131886) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(anim_from.AsCTypePtr()), gdc.ConstTypePtr(anim_to.AsCTypePtr()), gdc.ConstTypePtr(&sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetBlendTime(anim_from StringName, anim_to StringName, ) float32 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1958752504) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(anim_from.AsCTypePtr()), gdc.ConstTypePtr(anim_to.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetDefaultBlendTime(sec float32, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_blend_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetDefaultBlendTime() float32 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_blend_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) Play(name StringName, custom_blend float32, custom_speed float32, from_end bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3118260607) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&custom_blend), gdc.ConstTypePtr(&custom_speed), gdc.ConstTypePtr(&from_end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) PlayBackwards(name StringName, custom_blend float32, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_backwards")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2787282401) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&custom_blend), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) Pause()  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pause")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) Stop(keep_state bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 107499316) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep_state), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) IsPlaying() bool {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetCurrentAnimation(anim String, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(anim.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetCurrentAnimation() String {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetAssignedAnimation(anim String, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_assigned_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(anim.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetAssignedAnimation() String {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_assigned_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) Queue(name StringName, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetQueue() PackedStringArray {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_queue")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) ClearQueue()  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_queue")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) SetActive(active bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) IsActive() bool {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetSpeedScale(speed float32, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetSpeedScale() float32 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) GetPlayingSpeed() float32 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playing_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetAutoplay(name String, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetAutoplay() String {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetResetOnSaveEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reset_on_save_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) IsResetOnSaveEnabled() bool {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_reset_on_save_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetRoot(path NodePath, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetRoot() NodePath {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) FindAnimation(animation Animation, ) StringName {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1559484580) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(animation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) FindAnimationLibrary(animation Animation, ) StringName {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_animation_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1559484580) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(animation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) ClearCaches()  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_caches")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) SetProcessCallback(mode AnimationPlayerAnimationProcessCallback, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1663839457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetProcessCallback() AnimationPlayerAnimationProcessCallback {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4207496604) // FIXME: should cache?
  var ret AnimationPlayerAnimationProcessCallback
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetMethodCallMode(mode AnimationPlayerAnimationMethodCallMode, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_method_call_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3413514846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetMethodCallMode() AnimationPlayerAnimationMethodCallMode {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_method_call_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3583380054) // FIXME: should cache?
  var ret AnimationPlayerAnimationMethodCallMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetAudioMaxPolyphony(max_polyphony int, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_audio_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_polyphony), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) GetAudioMaxPolyphony() int {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_audio_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) SetMovieQuitOnFinishEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_movie_quit_on_finish_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) IsMovieQuitOnFinishEnabled() bool {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_movie_quit_on_finish_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) GetCurrentAnimationPosition() float32 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_animation_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) GetCurrentAnimationLength() float32 {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_animation_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationPlayer) Seek(seconds float32, update bool, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("seek")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2087892650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds), gdc.ConstTypePtr(&update), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationPlayer) Advance(delta float32, )  {
  classNameV := StringNameFromStr("AnimationPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("advance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationPlayerAnimationFinishedSignalFn func(anim_name StringName, )

func (me *AnimationPlayer) ConnectAnimationFinished(subs SignalSubscribers, fn AnimationPlayerAnimationFinishedSignalFn) {
  sig := StringNameFromStr("animation_finished")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectAnimationFinished(subs SignalSubscribers, fn AnimationPlayerAnimationFinishedSignalFn) {
  sig := StringNameFromStr("animation_finished")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type AnimationPlayerAnimationChangedSignalFn func(old_name StringName, new_name StringName, )

func (me *AnimationPlayer) ConnectAnimationChanged(subs SignalSubscribers, fn AnimationPlayerAnimationChangedSignalFn) {
  sig := StringNameFromStr("animation_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectAnimationChanged(subs SignalSubscribers, fn AnimationPlayerAnimationChangedSignalFn) {
  sig := StringNameFromStr("animation_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type AnimationPlayerAnimationStartedSignalFn func(anim_name StringName, )

func (me *AnimationPlayer) ConnectAnimationStarted(subs SignalSubscribers, fn AnimationPlayerAnimationStartedSignalFn) {
  sig := StringNameFromStr("animation_started")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectAnimationStarted(subs SignalSubscribers, fn AnimationPlayerAnimationStartedSignalFn) {
  sig := StringNameFromStr("animation_started")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type AnimationPlayerAnimationListChangedSignalFn func()

func (me *AnimationPlayer) ConnectAnimationListChanged(subs SignalSubscribers, fn AnimationPlayerAnimationListChangedSignalFn) {
  sig := StringNameFromStr("animation_list_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectAnimationListChanged(subs SignalSubscribers, fn AnimationPlayerAnimationListChangedSignalFn) {
  sig := StringNameFromStr("animation_list_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type AnimationPlayerAnimationLibrariesUpdatedSignalFn func()

func (me *AnimationPlayer) ConnectAnimationLibrariesUpdated(subs SignalSubscribers, fn AnimationPlayerAnimationLibrariesUpdatedSignalFn) {
  sig := StringNameFromStr("animation_libraries_updated")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectAnimationLibrariesUpdated(subs SignalSubscribers, fn AnimationPlayerAnimationLibrariesUpdatedSignalFn) {
  sig := StringNameFromStr("animation_libraries_updated")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type AnimationPlayerCachesClearedSignalFn func()

func (me *AnimationPlayer) ConnectCachesCleared(subs SignalSubscribers, fn AnimationPlayerCachesClearedSignalFn) {
  sig := StringNameFromStr("caches_cleared")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectCachesCleared(subs SignalSubscribers, fn AnimationPlayerCachesClearedSignalFn) {
  sig := StringNameFromStr("caches_cleared")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

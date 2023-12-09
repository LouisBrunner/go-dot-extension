// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *AnimationPlayer) XPostProcessKeyValue(animation Animation, track int, value Variant, object Object, object_idx int, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) AddAnimationLibrary(name StringName, library AnimationLibrary, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) RemoveAnimationLibrary(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) RenameAnimationLibrary(name StringName, newname StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) HasAnimationLibrary(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetAnimationLibrary(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetAnimationLibraryList()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) HasAnimation(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetAnimation(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetAnimationList()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) AnimationSetNext(anim_from StringName, anim_to StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) AnimationGetNext(anim_from StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetBlendTime(anim_from StringName, anim_to StringName, sec float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetBlendTime(anim_from StringName, anim_to StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetDefaultBlendTime(sec float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetDefaultBlendTime()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) Play(name StringName, custom_blend float32, custom_speed float32, from_end bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) PlayBackwards(name StringName, custom_blend float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) Pause()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) Stop(keep_state bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) IsPlaying()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetCurrentAnimation(anim String, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetCurrentAnimation()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetAssignedAnimation(anim String, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetAssignedAnimation()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) Queue(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetQueue()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) ClearQueue()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetActive(active bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) IsActive()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetSpeedScale(speed float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetSpeedScale()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetPlayingSpeed()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetAutoplay(name String, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetAutoplay()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetResetOnSaveEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) IsResetOnSaveEnabled()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetRoot(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetRoot()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) FindAnimation(animation Animation, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) FindAnimationLibrary(animation Animation, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) ClearCaches()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetProcessCallback(mode AnimationPlayerAnimationProcessCallback, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetProcessCallback()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetMethodCallMode(mode AnimationPlayerAnimationMethodCallMode, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetMethodCallMode()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetAudioMaxPolyphony(max_polyphony int, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetAudioMaxPolyphony()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) SetMovieQuitOnFinishEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) IsMovieQuitOnFinishEnabled()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetCurrentAnimationPosition()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) GetCurrentAnimationLength()  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) Seek(seconds float32, update bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationPlayer) Advance(delta float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

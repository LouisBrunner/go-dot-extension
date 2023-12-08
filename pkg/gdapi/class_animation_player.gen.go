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

func  (me *AnimationPlayer) XPostProcessKeyValue(animation Animation, track int, value Variant, object Object, object_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) AddAnimationLibrary(name StringName, library AnimationLibrary, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) RemoveAnimationLibrary(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) RenameAnimationLibrary(name StringName, newname StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) HasAnimationLibrary(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetAnimationLibrary(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetAnimationLibraryList() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) HasAnimation(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetAnimation(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetAnimationList() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) AnimationSetNext(anim_from StringName, anim_to StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) AnimationGetNext(anim_from StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetBlendTime(anim_from StringName, anim_to StringName, sec float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetBlendTime(anim_from StringName, anim_to StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetDefaultBlendTime(sec float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetDefaultBlendTime() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) Play(name StringName, custom_blend float32, custom_speed float32, from_end bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) PlayBackwards(name StringName, custom_blend float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) Pause() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) Stop(keep_state bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) IsPlaying() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetCurrentAnimation(anim String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetCurrentAnimation() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetAssignedAnimation(anim String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetAssignedAnimation() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) Queue(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetQueue() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) ClearQueue() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetActive(active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) IsActive() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetSpeedScale(speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetSpeedScale() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetPlayingSpeed() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetAutoplay(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetAutoplay() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetResetOnSaveEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) IsResetOnSaveEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetRoot(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetRoot() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) FindAnimation(animation Animation, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) FindAnimationLibrary(animation Animation, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) ClearCaches() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetProcessCallback(mode AnimationPlayerAnimationProcessCallback, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetProcessCallback() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetMethodCallMode(mode AnimationPlayerAnimationMethodCallMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetMethodCallMode() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetAudioMaxPolyphony(max_polyphony int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetAudioMaxPolyphony() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) SetMovieQuitOnFinishEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) IsMovieQuitOnFinishEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetCurrentAnimationPosition() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) GetCurrentAnimationLength() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) Seek(seconds float32, update bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationPlayer) Advance(delta float32, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals

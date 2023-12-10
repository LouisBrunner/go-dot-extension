// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimatedSprite3D struct {
  obj gdc.ObjectPtr
}

func (me *AnimatedSprite3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatedSprite3D) BaseClass() string {
  return "AnimatedSprite3D"
}



// Enums

func (me *AnimatedSprite3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimatedSprite3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimatedSprite3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimatedSprite3D) SetSpriteFrames(sprite_frames SpriteFrames, )  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sprite_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 905781144) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(sprite_frames.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) GetSpriteFrames() SpriteFrames {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sprite_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3804851214) // FIXME: should cache?
  var ret SpriteFrames
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite3D) SetAnimation(name StringName, )  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) GetAnimation() StringName {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite3D) SetAutoplay(name String, )  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) GetAutoplay() String {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite3D) IsPlaying() bool {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite3D) Play(name StringName, custom_speed float32, from_end bool, )  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2372066587) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&custom_speed), gdc.ConstTypePtr(&from_end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) PlayBackwards(name StringName, )  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_backwards")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1421762485) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) Pause()  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pause")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) Stop()  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) SetFrame(frame int, )  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) GetFrame() int {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite3D) SetFrameProgress(progress float32, )  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&progress), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) GetFrameProgress() float32 {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite3D) SetFrameAndProgress(frame int, progress float32, )  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame_and_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame), gdc.ConstTypePtr(&progress), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) SetSpeedScale(speed_scale float32, )  {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite3D) GetSpeedScale() float32 {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite3D) GetPlayingSpeed() float32 {
  classNameV := StringNameFromStr("AnimatedSprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playing_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AnimatedSprite3D) GetPropSpriteFrames() SpriteFrames {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) SetPropSpriteFrames(value SpriteFrames) {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) GetPropAnimation() StringName {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) SetPropAnimation(value StringName) {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) GetPropAutoplay() StringName {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) SetPropAutoplay(value StringName) {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) GetPropFrame() int {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) SetPropFrame(value int) {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) GetPropFrameProgress() float32 {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) SetPropFrameProgress(value float32) {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) GetPropSpeedScale() float32 {
  panic("TODO: implement")
}

func (me *AnimatedSprite3D) SetPropSpeedScale(value float32) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API
// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimatedSprite2D struct {
  obj gdc.ObjectPtr
}

func (me *AnimatedSprite2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatedSprite2D) BaseClass() string {
  return "AnimatedSprite2D"
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
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sprite_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 905781144) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(sprite_frames.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) GetSpriteFrames() SpriteFrames {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sprite_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3804851214) // FIXME: should cache?
  var ret SpriteFrames
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) SetAnimation(name StringName, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) GetAnimation() StringName {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) SetAutoplay(name String, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) GetAutoplay() String {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) IsPlaying() bool {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) Play(name StringName, custom_speed float32, from_end bool, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2372066587) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&custom_speed), gdc.ConstTypePtr(&from_end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) PlayBackwards(name StringName, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_backwards")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1421762485) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) Pause()  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pause")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) Stop()  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) SetCentered(centered bool, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&centered), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) IsCentered() bool {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) GetOffset() Vector2 {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) SetFlipH(flip_h bool, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_h), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) IsFlippedH() bool {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) SetFlipV(flip_v bool, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_v), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) IsFlippedV() bool {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) SetFrame(frame int, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) GetFrame() int {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) SetFrameProgress(progress float32, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&progress), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) GetFrameProgress() float32 {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) SetFrameAndProgress(frame int, progress float32, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame_and_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame), gdc.ConstTypePtr(&progress), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) SetSpeedScale(speed_scale float32, )  {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatedSprite2D) GetSpeedScale() float32 {
  classNameV := StringNameFromStr("AnimatedSprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimatedSprite2D) GetPlayingSpeed() float32 {
  classNameV := StringNameFromStr("AnimatedSprite2D")
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

func (me *AnimatedSprite2D) GetPropSpriteFrames() SpriteFrames {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropSpriteFrames(value SpriteFrames) {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) GetPropAnimation() StringName {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropAnimation(value StringName) {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) GetPropAutoplay() StringName {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropAutoplay(value StringName) {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) GetPropFrame() int {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropFrame(value int) {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) GetPropFrameProgress() float32 {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropFrameProgress(value float32) {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) GetPropSpeedScale() float32 {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropSpeedScale(value float32) {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) GetPropCentered() bool {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropCentered(value bool) {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) GetPropOffset() Vector2 {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) GetPropFlipH() bool {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropFlipH(value bool) {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) GetPropFlipV() bool {
  panic("TODO: implement")
}

func (me *AnimatedSprite2D) SetPropFlipV(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API
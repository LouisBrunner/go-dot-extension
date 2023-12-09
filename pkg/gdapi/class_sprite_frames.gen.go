// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpriteFrames struct {
  obj gdc.ObjectPtr
}

func (me *SpriteFrames) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SpriteFrames) BaseClass() string {
  return "SpriteFrames"
}



// Enums

func (me *SpriteFrames) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpriteFrames) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SpriteFrames) AddAnimation(anim StringName, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) HasAnimation(anim StringName, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) RemoveAnimation(anim StringName, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) RenameAnimation(anim StringName, newname StringName, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) GetAnimationNames()  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) SetAnimationSpeed(anim StringName, fps float32, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) GetAnimationSpeed(anim StringName, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) SetAnimationLoop(anim StringName, loop bool, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) GetAnimationLoop(anim StringName, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) AddFrame(anim StringName, texture Texture2D, duration float32, at_position int, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) SetFrame(anim StringName, idx int, texture Texture2D, duration float32, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) RemoveFrame(anim StringName, idx int, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) GetFrameCount(anim StringName, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) GetFrameTexture(anim StringName, idx int, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) GetFrameDuration(anim StringName, idx int, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) Clear(anim StringName, )  {
  panic("TODO: implement")
}

func  (me *SpriteFrames) ClearAll()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

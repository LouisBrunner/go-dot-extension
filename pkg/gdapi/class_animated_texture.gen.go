// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatedTexture struct {
  obj gdc.ObjectPtr
}

func (me *AnimatedTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatedTexture) BaseClass() string {
  return "AnimatedTexture"
}



// Constants

var (
  AnimatedTextureMaxFrames = "256" // TODO: construct correctly
)

// Enums

func (me *AnimatedTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimatedTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimatedTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimatedTexture) SetFrames(frames int, )  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) GetFrames()  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) SetCurrentFrame(frame int, )  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) GetCurrentFrame()  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) SetPause(pause bool, )  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) GetPause()  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) SetOneShot(one_shot bool, )  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) GetOneShot()  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) SetSpeedScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) GetSpeedScale()  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) SetFrameTexture(frame int, texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) GetFrameTexture(frame int, )  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) SetFrameDuration(frame int, duration float32, )  {
  panic("TODO: implement")
}

func  (me *AnimatedTexture) GetFrameDuration(frame int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

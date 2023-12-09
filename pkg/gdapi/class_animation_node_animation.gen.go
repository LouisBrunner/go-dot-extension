// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeAnimation struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeAnimation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeAnimation) BaseClass() string {
  return "AnimationNodeAnimation"
}



// Enums

type AnimationNodeAnimationPlayMode int
const (
  AnimationNodeAnimationPlayModePlayModeForward AnimationNodeAnimationPlayMode = 0
  AnimationNodeAnimationPlayModePlayModeBackward AnimationNodeAnimationPlayMode = 1
)

func (me *AnimationNodeAnimation) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeAnimation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeAnimation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationNodeAnimation) SetAnimation(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeAnimation) GetAnimation()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeAnimation) SetPlayMode(mode AnimationNodeAnimationPlayMode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeAnimation) GetPlayMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

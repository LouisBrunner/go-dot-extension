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

type AnimationNodeAnimationPlayMode int
const (
  AnimationNodeAnimationPlayModePlayModeForward AnimationNodeAnimationPlayMode = 0
  AnimationNodeAnimationPlayModePlayModeBackward AnimationNodeAnimationPlayMode = 1
)

func  (me *AnimationNodeAnimation) SetAnimation(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeAnimation) GetAnimation() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeAnimation) SetPlayMode(mode AnimationNodeAnimationPlayMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeAnimation) GetPlayMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals

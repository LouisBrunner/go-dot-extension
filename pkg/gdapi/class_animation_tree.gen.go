// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationTree struct {
  obj gdc.ObjectPtr
}

func (me *AnimationTree) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationTree) BaseClass() string {
  return "AnimationTree"
}

type AnimationTreeAnimationProcessCallback int
const (
  AnimationTreeAnimationProcessPhysics AnimationTreeAnimationProcessCallback = 0
  AnimationTreeAnimationProcessIdle AnimationTreeAnimationProcessCallback = 1
  AnimationTreeAnimationProcessManual AnimationTreeAnimationProcessCallback = 2
)

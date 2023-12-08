// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeOneShot struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeOneShot) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeOneShot) BaseClass() string {
  return "AnimationNodeOneShot"
}

type AnimationNodeOneShotOneShotRequest int
const (
  AnimationNodeOneShotOneShotRequestNone AnimationNodeOneShotOneShotRequest = 0
  AnimationNodeOneShotOneShotRequestFire AnimationNodeOneShotOneShotRequest = 1
  AnimationNodeOneShotOneShotRequestAbort AnimationNodeOneShotOneShotRequest = 2
  AnimationNodeOneShotOneShotRequestFadeOut AnimationNodeOneShotOneShotRequest = 3
)

type AnimationNodeOneShotMixMode int
const (
  AnimationNodeOneShotMixModeBlend AnimationNodeOneShotMixMode = 0
  AnimationNodeOneShotMixModeAdd AnimationNodeOneShotMixMode = 1
)

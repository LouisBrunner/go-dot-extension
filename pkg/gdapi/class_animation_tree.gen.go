// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  AnimationTreeAnimationProcessCallbackAnimationProcessPhysics AnimationTreeAnimationProcessCallback = 0
  AnimationTreeAnimationProcessCallbackAnimationProcessIdle AnimationTreeAnimationProcessCallback = 1
  AnimationTreeAnimationProcessCallbackAnimationProcessManual AnimationTreeAnimationProcessCallback = 2
)

func  (me *AnimationTree) XPostProcessKeyValue(animation Animation, track int, value Variant, object Object, object_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) SetActive(active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) IsActive() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) SetTreeRoot(root AnimationNode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetTreeRoot() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) SetProcessCallback(mode AnimationTreeAnimationProcessCallback, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetProcessCallback() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) SetAnimationPlayer(root NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetAnimationPlayer() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) SetAdvanceExpressionBaseNode(node NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetAdvanceExpressionBaseNode() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) SetRootMotionTrack(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetRootMotionTrack() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) SetAudioMaxPolyphony(max_polyphony int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetAudioMaxPolyphony() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetRootMotionPosition() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetRootMotionRotation() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetRootMotionScale() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetRootMotionPositionAccumulator() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetRootMotionRotationAccumulator() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) GetRootMotionScaleAccumulator() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationTree) Advance(delta float32, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals

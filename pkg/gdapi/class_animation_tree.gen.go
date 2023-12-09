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



// Enums

type AnimationTreeAnimationProcessCallback int
const (
  AnimationTreeAnimationProcessCallbackAnimationProcessPhysics AnimationTreeAnimationProcessCallback = 0
  AnimationTreeAnimationProcessCallbackAnimationProcessIdle AnimationTreeAnimationProcessCallback = 1
  AnimationTreeAnimationProcessCallbackAnimationProcessManual AnimationTreeAnimationProcessCallback = 2
)

func (me *AnimationTree) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationTree) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationTree) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationTree) XPostProcessKeyValue(animation Animation, track int, value Variant, object Object, object_idx int, )  {
  panic("TODO: implement")
}

func  (me *AnimationTree) SetActive(active bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationTree) IsActive()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) SetTreeRoot(root AnimationNode, )  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetTreeRoot()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) SetProcessCallback(mode AnimationTreeAnimationProcessCallback, )  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetProcessCallback()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) SetAnimationPlayer(root NodePath, )  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetAnimationPlayer()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) SetAdvanceExpressionBaseNode(node NodePath, )  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetAdvanceExpressionBaseNode()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) SetRootMotionTrack(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetRootMotionTrack()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) SetAudioMaxPolyphony(max_polyphony int, )  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetAudioMaxPolyphony()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetRootMotionPosition()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetRootMotionRotation()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetRootMotionScale()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetRootMotionPositionAccumulator()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetRootMotionRotationAccumulator()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) GetRootMotionScaleAccumulator()  {
  panic("TODO: implement")
}

func  (me *AnimationTree) Advance(delta float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

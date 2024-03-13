// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OpenXRHand struct {
  obj gdc.ObjectPtr
}

func (me *OpenXRHand) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OpenXRHand) BaseClass() string {
  return "OpenXRHand"
}



// Enums

type OpenXRHandHands int
const (
  OpenXRHandHandsHandLeft OpenXRHandHands = 0
  OpenXRHandHandsHandRight OpenXRHandHands = 1
  OpenXRHandHandsHandMax OpenXRHandHands = 2
)

type OpenXRHandMotionRange int
const (
  OpenXRHandMotionRangeMotionRangeUnobstructed OpenXRHandMotionRange = 0
  OpenXRHandMotionRangeMotionRangeConformToController OpenXRHandMotionRange = 1
  OpenXRHandMotionRangeMotionRangeMax OpenXRHandMotionRange = 2
)

func (me *OpenXRHand) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRHand) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRHand) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRHand) SetHand(hand OpenXRHandHands, )  {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1849328560) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRHand) GetHand() OpenXRHandHands {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2850644561) // FIXME: should cache?
  var ret OpenXRHandHands
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRHand) SetHandSkeleton(hand_skeleton NodePath, )  {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hand_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(hand_skeleton.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRHand) GetHandSkeleton() NodePath {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRHand) SetMotionRange(motion_range OpenXRHandMotionRange, )  {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3326516003) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&motion_range), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRHand) GetMotionRange() OpenXRHandMotionRange {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2191822314) // FIXME: should cache?
  var ret OpenXRHandMotionRange
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

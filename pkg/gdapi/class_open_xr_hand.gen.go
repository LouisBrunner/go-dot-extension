// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type OpenXRHand struct {
  Node3D
}

func (me *OpenXRHand) BaseClass() string {
  return "OpenXRHand"
}

func NewOpenXRHand() *OpenXRHand {
  str := StringNameFromStr("OpenXRHand") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OpenXRHand{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRHand) GetHand() OpenXRHandHands {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2850644561) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret OpenXRHandHands

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *OpenXRHand) SetHandSkeleton(hand_skeleton NodePath, )  {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hand_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{hand_skeleton.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRHand) GetHandSkeleton() NodePath {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRHand) SetMotionRange(motion_range OpenXRHandMotionRange, )  {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3326516003) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&motion_range) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRHand) GetMotionRange() OpenXRHandMotionRange {
  classNameV := StringNameFromStr("OpenXRHand")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2191822314) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret OpenXRHandMotionRange

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForOpenXRHandList struct {
  fnSetHand gdc.MethodBindPtr
  fnGetHand gdc.MethodBindPtr
  fnSetHandSkeleton gdc.MethodBindPtr
  fnGetHandSkeleton gdc.MethodBindPtr
  fnSetMotionRange gdc.MethodBindPtr
  fnGetMotionRange gdc.MethodBindPtr
}

var ptrsForOpenXRHand ptrsForOpenXRHandList

func initOpenXRHandPtrs(iface gdc.Interface) {

  className := StringNameFromStr("OpenXRHand")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_hand")
    defer methodName.Destroy()
    ptrsForOpenXRHand.fnSetHand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1849328560))
  }
  {
    methodName := StringNameFromStr("get_hand")
    defer methodName.Destroy()
    ptrsForOpenXRHand.fnGetHand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2850644561))
  }
  {
    methodName := StringNameFromStr("set_hand_skeleton")
    defer methodName.Destroy()
    ptrsForOpenXRHand.fnSetHandSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_hand_skeleton")
    defer methodName.Destroy()
    ptrsForOpenXRHand.fnGetHandSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_motion_range")
    defer methodName.Destroy()
    ptrsForOpenXRHand.fnSetMotionRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3326516003))
  }
  {
    methodName := StringNameFromStr("get_motion_range")
    defer methodName.Destroy()
    ptrsForOpenXRHand.fnGetMotionRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2191822314))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRHand.fnSetHand), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRHand) GetHand() OpenXRHandHands {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret OpenXRHandHands

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRHand.fnGetHand), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *OpenXRHand) SetHandSkeleton(hand_skeleton NodePath, )  {
  cargs := []gdc.ConstTypePtr{hand_skeleton.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRHand.fnSetHandSkeleton), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRHand) GetHandSkeleton() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRHand.fnGetHandSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRHand) SetMotionRange(motion_range OpenXRHandMotionRange, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&motion_range) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRHand.fnSetMotionRange), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRHand) GetMotionRange() OpenXRHandMotionRange {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret OpenXRHandMotionRange

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRHand.fnGetMotionRange), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

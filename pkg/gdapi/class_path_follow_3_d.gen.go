// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PathFollow3D struct {
  obj gdc.ObjectPtr
}

func (me *PathFollow3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PathFollow3D) BaseClass() string {
  return "PathFollow3D"
}



// Enums

type PathFollow3DRotationMode int
const (
  PathFollow3DRotationModeRotationNone PathFollow3DRotationMode = 0
  PathFollow3DRotationModeRotationY PathFollow3DRotationMode = 1
  PathFollow3DRotationModeRotationXy PathFollow3DRotationMode = 2
  PathFollow3DRotationModeRotationXyz PathFollow3DRotationMode = 3
  PathFollow3DRotationModeRotationOriented PathFollow3DRotationMode = 4
)

func (me *PathFollow3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PathFollow3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PathFollow3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PathFollow3D) SetProgress(progress float32, )  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) GetProgress()  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) SetHOffset(h_offset float32, )  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) GetHOffset()  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) SetVOffset(v_offset float32, )  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) GetVOffset()  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) SetProgressRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) GetProgressRatio()  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) SetRotationMode(rotation_mode PathFollow3DRotationMode, )  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) GetRotationMode()  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) SetCubicInterpolation(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) GetCubicInterpolation()  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) SetUseModelFront(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) IsUsingModelFront()  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) SetLoop(loop bool, )  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) HasLoop()  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) SetTiltEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PathFollow3D) IsTiltEnabled()  {
  panic("TODO: implement")
}

func  PathFollow3DCorrectPosture(transform Transform3D, rotation_mode PathFollow3DRotationMode, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

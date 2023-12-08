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

type PathFollow3DRotationMode int
const (
  PathFollow3DRotationModeRotationNone PathFollow3DRotationMode = 0
  PathFollow3DRotationModeRotationY PathFollow3DRotationMode = 1
  PathFollow3DRotationModeRotationXy PathFollow3DRotationMode = 2
  PathFollow3DRotationModeRotationXyz PathFollow3DRotationMode = 3
  PathFollow3DRotationModeRotationOriented PathFollow3DRotationMode = 4
)

func  (me *PathFollow3D) SetProgress(progress float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) GetProgress() { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) SetHOffset(h_offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) GetHOffset() { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) SetVOffset(v_offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) GetVOffset() { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) SetProgressRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) GetProgressRatio() { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) SetRotationMode(rotation_mode PathFollow3DRotationMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) GetRotationMode() { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) SetCubicInterpolation(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) GetCubicInterpolation() { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) SetUseModelFront(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) IsUsingModelFront() { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) SetLoop(loop bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) HasLoop() { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) SetTiltEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PathFollow3D) IsTiltEnabled() { // TODO: return value
  // TODO: implement
}

func  PathFollow3DCorrectPosture(transform Transform3D, rotation_mode PathFollow3DRotationMode, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Gradient struct {
  obj gdc.ObjectPtr
}

func (me *Gradient) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Gradient) BaseClass() string {
  return "Gradient"
}

type GradientInterpolationMode int
const (
  GradientInterpolationModeGradientInterpolateLinear GradientInterpolationMode = 0
  GradientInterpolationModeGradientInterpolateConstant GradientInterpolationMode = 1
  GradientInterpolationModeGradientInterpolateCubic GradientInterpolationMode = 2
)

type GradientColorSpace int
const (
  GradientColorSpaceGradientColorSpaceSrgb GradientColorSpace = 0
  GradientColorSpaceGradientColorSpaceLinearSrgb GradientColorSpace = 1
  GradientColorSpaceGradientColorSpaceOklab GradientColorSpace = 2
)

func  (me *Gradient) AddPoint(offset float32, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) RemovePoint(point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) SetOffset(point int, offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) GetOffset(point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) Reverse() { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) SetColor(point int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) GetColor(point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) Sample(offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) GetPointCount() { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) SetOffsets(offsets PackedFloat32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) GetOffsets() { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) SetColors(colors PackedColorArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) GetColors() { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) SetInterpolationMode(interpolation_mode GradientInterpolationMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) GetInterpolationMode() { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) SetInterpolationColorSpace(interpolation_color_space GradientColorSpace, ) { // TODO: return value
  // TODO: implement
}

func  (me *Gradient) GetInterpolationColorSpace() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals

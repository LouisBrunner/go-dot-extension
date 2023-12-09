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



// Enums

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

func (me *Gradient) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Gradient) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Gradient) AddPoint(offset float32, color Color, )  {
  panic("TODO: implement")
}

func  (me *Gradient) RemovePoint(point int, )  {
  panic("TODO: implement")
}

func  (me *Gradient) SetOffset(point int, offset float32, )  {
  panic("TODO: implement")
}

func  (me *Gradient) GetOffset(point int, )  {
  panic("TODO: implement")
}

func  (me *Gradient) Reverse()  {
  panic("TODO: implement")
}

func  (me *Gradient) SetColor(point int, color Color, )  {
  panic("TODO: implement")
}

func  (me *Gradient) GetColor(point int, )  {
  panic("TODO: implement")
}

func  (me *Gradient) Sample(offset float32, )  {
  panic("TODO: implement")
}

func  (me *Gradient) GetPointCount()  {
  panic("TODO: implement")
}

func  (me *Gradient) SetOffsets(offsets PackedFloat32Array, )  {
  panic("TODO: implement")
}

func  (me *Gradient) GetOffsets()  {
  panic("TODO: implement")
}

func  (me *Gradient) SetColors(colors PackedColorArray, )  {
  panic("TODO: implement")
}

func  (me *Gradient) GetColors()  {
  panic("TODO: implement")
}

func  (me *Gradient) SetInterpolationMode(interpolation_mode GradientInterpolationMode, )  {
  panic("TODO: implement")
}

func  (me *Gradient) GetInterpolationMode()  {
  panic("TODO: implement")
}

func  (me *Gradient) SetInterpolationColorSpace(interpolation_color_space GradientColorSpace, )  {
  panic("TODO: implement")
}

func  (me *Gradient) GetInterpolationColorSpace()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

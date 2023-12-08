// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  GradientGradientInterpolateLinear GradientInterpolationMode = 0
  GradientGradientInterpolateConstant GradientInterpolationMode = 1
  GradientGradientInterpolateCubic GradientInterpolationMode = 2
)

type GradientColorSpace int
const (
  GradientGradientColorSpaceSrgb GradientColorSpace = 0
  GradientGradientColorSpaceLinearSrgb GradientColorSpace = 1
  GradientGradientColorSpaceOklab GradientColorSpace = 2
)

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve struct {
  obj gdc.ObjectPtr
}

func (me *Curve) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Curve) BaseClass() string {
  return "Curve"
}

type CurveTangentMode int
const (
  CurveTangentModeTangentFree CurveTangentMode = 0
  CurveTangentModeTangentLinear CurveTangentMode = 1
  CurveTangentModeTangentModeCount CurveTangentMode = 2
)

func  (me *Curve) GetPointCount() { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetPointCount(count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) AddPoint(position Vector2, left_tangent float32, right_tangent float32, left_mode CurveTangentMode, right_mode CurveTangentMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) RemovePoint(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) ClearPoints() { // TODO: return value
  // TODO: implement
}

func  (me *Curve) GetPointPosition(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetPointValue(index int, y float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetPointOffset(index int, offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) Sample(offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SampleBaked(offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) GetPointLeftTangent(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) GetPointRightTangent(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) GetPointLeftMode(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) GetPointRightMode(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetPointLeftTangent(index int, tangent float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetPointRightTangent(index int, tangent float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetPointLeftMode(index int, mode CurveTangentMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetPointRightMode(index int, mode CurveTangentMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) GetMinValue() { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetMinValue(min float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) GetMaxValue() { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetMaxValue(max float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Curve) CleanDupes() { // TODO: return value
  // TODO: implement
}

func  (me *Curve) Bake() { // TODO: return value
  // TODO: implement
}

func  (me *Curve) GetBakeResolution() { // TODO: return value
  // TODO: implement
}

func  (me *Curve) SetBakeResolution(resolution int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals

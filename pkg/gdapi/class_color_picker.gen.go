// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ColorPicker struct {
  obj gdc.ObjectPtr
}

func (me *ColorPicker) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ColorPicker) BaseClass() string {
  return "ColorPicker"
}

type ColorPickerColorModeType int
const (
  ColorPickerModeRgb ColorPickerColorModeType = 0
  ColorPickerModeHsv ColorPickerColorModeType = 1
  ColorPickerModeRaw ColorPickerColorModeType = 2
  ColorPickerModeOkhsl ColorPickerColorModeType = 3
)

type ColorPickerPickerShapeType int
const (
  ColorPickerShapeHsvRectangle ColorPickerPickerShapeType = 0
  ColorPickerShapeHsvWheel ColorPickerPickerShapeType = 1
  ColorPickerShapeVhsCircle ColorPickerPickerShapeType = 2
  ColorPickerShapeOkhslCircle ColorPickerPickerShapeType = 3
  ColorPickerShapeNone ColorPickerPickerShapeType = 4
)

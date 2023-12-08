// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Input struct {
  obj gdc.ObjectPtr
}

func (me *Input) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Input) BaseClass() string {
  return "Input"
}

type InputMouseMode int
const (
  InputMouseModeVisible InputMouseMode = 0
  InputMouseModeHidden InputMouseMode = 1
  InputMouseModeCaptured InputMouseMode = 2
  InputMouseModeConfined InputMouseMode = 3
  InputMouseModeConfinedHidden InputMouseMode = 4
)

type InputCursorShape int
const (
  InputCursorArrow InputCursorShape = 0
  InputCursorIbeam InputCursorShape = 1
  InputCursorPointingHand InputCursorShape = 2
  InputCursorCross InputCursorShape = 3
  InputCursorWait InputCursorShape = 4
  InputCursorBusy InputCursorShape = 5
  InputCursorDrag InputCursorShape = 6
  InputCursorCanDrop InputCursorShape = 7
  InputCursorForbidden InputCursorShape = 8
  InputCursorVsize InputCursorShape = 9
  InputCursorHsize InputCursorShape = 10
  InputCursorBdiagsize InputCursorShape = 11
  InputCursorFdiagsize InputCursorShape = 12
  InputCursorMove InputCursorShape = 13
  InputCursorVsplit InputCursorShape = 14
  InputCursorHsplit InputCursorShape = 15
  InputCursorHelp InputCursorShape = 16
)

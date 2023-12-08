// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TouchScreenButton struct {
  obj gdc.ObjectPtr
}

func (me *TouchScreenButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TouchScreenButton) BaseClass() string {
  return "TouchScreenButton"
}

type TouchScreenButtonVisibilityMode int
const (
  TouchScreenButtonVisibilityAlways TouchScreenButtonVisibilityMode = 0
  TouchScreenButtonVisibilityTouchscreenOnly TouchScreenButtonVisibilityMode = 1
)

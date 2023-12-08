// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RootMotionView struct {
  obj gdc.ObjectPtr
}

func createRootMotionView(obj gdc.ObjectPtr) *RootMotionView {
  return &RootMotionView{
    obj: obj,
  }
}

func (me *RootMotionView) BaseClass() string {
  return "RootMotionView"
}

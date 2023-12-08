// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RenderingDevice struct {
  obj gdc.ObjectPtr
}

func createRenderingDevice(obj gdc.ObjectPtr) *RenderingDevice {
  return &RenderingDevice{
    obj: obj,
  }
}

func (me *RenderingDevice) BaseClass() string {
  return "RenderingDevice"
}

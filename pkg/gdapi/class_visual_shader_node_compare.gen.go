// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeCompare struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeCompare(obj gdc.ObjectPtr) *VisualShaderNodeCompare {
  return &VisualShaderNodeCompare{
    obj: obj,
  }
}

func (me *VisualShaderNodeCompare) BaseClass() string {
  return "VisualShaderNodeCompare"
}

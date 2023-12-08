// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShader struct {
  obj gdc.ObjectPtr
}

func createVisualShader(obj gdc.ObjectPtr) *VisualShader {
  return &VisualShader{
    obj: obj,
  }
}

func (me *VisualShader) BaseClass() string {
  return "VisualShader"
}

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeComment struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeComment(obj gdc.ObjectPtr) *VisualShaderNodeComment {
  return &VisualShaderNodeComment{
    obj: obj,
  }
}

func (me *VisualShaderNodeComment) BaseClass() string {
  return "VisualShaderNodeComment"
}

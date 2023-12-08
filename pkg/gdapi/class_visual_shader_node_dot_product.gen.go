// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeDotProduct struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeDotProduct(obj gdc.ObjectPtr) *VisualShaderNodeDotProduct {
  return &VisualShaderNodeDotProduct{
    obj: obj,
  }
}

func (me *VisualShaderNodeDotProduct) BaseClass() string {
  return "VisualShaderNodeDotProduct"
}

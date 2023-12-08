// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTexture2DArray struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTexture2DArray(obj gdc.ObjectPtr) *VisualShaderNodeTexture2DArray {
  return &VisualShaderNodeTexture2DArray{
    obj: obj,
  }
}

func (me *VisualShaderNodeTexture2DArray) BaseClass() string {
  return "VisualShaderNodeTexture2DArray"
}

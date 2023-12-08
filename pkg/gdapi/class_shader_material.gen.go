// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ShaderMaterial struct {
  obj gdc.ObjectPtr
}

func createShaderMaterial(obj gdc.ObjectPtr) *ShaderMaterial {
  return &ShaderMaterial{
    obj: obj,
  }
}

func (me *ShaderMaterial) BaseClass() string {
  return "ShaderMaterial"
}

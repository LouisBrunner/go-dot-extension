// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ShaderGlobalsOverride struct {
  obj gdc.ObjectPtr
}

func createShaderGlobalsOverride(obj gdc.ObjectPtr) *ShaderGlobalsOverride {
  return &ShaderGlobalsOverride{
    obj: obj,
  }
}

func (me *ShaderGlobalsOverride) BaseClass() string {
  return "ShaderGlobalsOverride"
}
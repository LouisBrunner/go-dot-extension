// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ShaderInclude struct {
  obj gdc.ObjectPtr
}

func createShaderInclude(obj gdc.ObjectPtr) *ShaderInclude {
  return &ShaderInclude{
    obj: obj,
  }
}

func (me *ShaderInclude) BaseClass() string {
  return "ShaderInclude"
}

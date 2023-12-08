// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shader struct {
  obj gdc.ObjectPtr
}

func createShader(obj gdc.ObjectPtr) *Shader {
  return &Shader{
    obj: obj,
  }
}

func (me *Shader) BaseClass() string {
  return "Shader"
}

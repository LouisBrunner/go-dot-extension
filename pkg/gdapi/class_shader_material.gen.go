// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ShaderMaterial struct {
  obj gdc.ObjectPtr
}

func (me *ShaderMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ShaderMaterial) BaseClass() string {
  return "ShaderMaterial"
}

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShader struct {
  obj gdc.ObjectPtr
}

func (me *VisualShader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShader) BaseClass() string {
  return "VisualShader"
}

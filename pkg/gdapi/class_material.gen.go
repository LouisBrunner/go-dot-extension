// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Material struct {
  obj gdc.ObjectPtr
}

func (me *Material) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Material) BaseClass() string {
  return "Material"
}

const (
  MaterialRENDER_PRIORITY_MAX = 127
  MaterialRENDER_PRIORITY_MIN = -128
)

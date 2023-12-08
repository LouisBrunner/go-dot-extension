// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Light3D struct {
  obj gdc.ObjectPtr
}

func (me *Light3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Light3D) BaseClass() string {
  return "Light3D"
}

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Light2D struct {
  obj gdc.ObjectPtr
}

func (me *Light2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Light2D) BaseClass() string {
  return "Light2D"
}

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CircleShape2D struct {
  obj gdc.ObjectPtr
}

func (me *CircleShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CircleShape2D) BaseClass() string {
  return "CircleShape2D"
}

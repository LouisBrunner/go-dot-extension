// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Tween struct {
  obj gdc.ObjectPtr
}

func createTween(obj gdc.ObjectPtr) *Tween {
  return &Tween{
    obj: obj,
  }
}

func (me *Tween) BaseClass() string {
  return "Tween"
}

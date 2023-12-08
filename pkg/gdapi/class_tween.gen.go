// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Tween struct {
  obj gdc.ObjectPtr
}

func (me *Tween) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Tween) BaseClass() string {
  return "Tween"
}

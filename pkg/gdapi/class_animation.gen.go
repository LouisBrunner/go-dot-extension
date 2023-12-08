// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Animation struct {
  obj gdc.ObjectPtr
}

func (me *Animation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Animation) BaseClass() string {
  return "Animation"
}

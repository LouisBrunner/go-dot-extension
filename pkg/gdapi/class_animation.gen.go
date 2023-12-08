// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Animation struct {
  obj gdc.ObjectPtr
}

func createAnimation(obj gdc.ObjectPtr) *Animation {
  return &Animation{
    obj: obj,
  }
}

func (me *Animation) BaseClass() string {
  return "Animation"
}

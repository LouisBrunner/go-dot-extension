// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Bone2D struct {
  obj gdc.ObjectPtr
}

func (me *Bone2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Bone2D) BaseClass() string {
  return "Bone2D"
}

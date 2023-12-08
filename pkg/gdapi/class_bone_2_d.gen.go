// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Bone2D struct {
  obj gdc.ObjectPtr
}

func createBone2D(obj gdc.ObjectPtr) *Bone2D {
  return &Bone2D{
    obj: obj,
  }
}

func (me *Bone2D) BaseClass() string {
  return "Bone2D"
}

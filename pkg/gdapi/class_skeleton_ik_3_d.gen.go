// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonIK3D struct {
  obj gdc.ObjectPtr
}

func createSkeletonIK3D(obj gdc.ObjectPtr) *SkeletonIK3D {
  return &SkeletonIK3D{
    obj: obj,
  }
}

func (me *SkeletonIK3D) BaseClass() string {
  return "SkeletonIK3D"
}

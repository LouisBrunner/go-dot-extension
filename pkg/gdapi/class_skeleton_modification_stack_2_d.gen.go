// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonModificationStack2D struct {
  obj gdc.ObjectPtr
}

func createSkeletonModificationStack2D(obj gdc.ObjectPtr) *SkeletonModificationStack2D {
  return &SkeletonModificationStack2D{
    obj: obj,
  }
}

func (me *SkeletonModificationStack2D) BaseClass() string {
  return "SkeletonModificationStack2D"
}

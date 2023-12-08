// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonModification2D struct {
  obj gdc.ObjectPtr
}

func createSkeletonModification2D(obj gdc.ObjectPtr) *SkeletonModification2D {
  return &SkeletonModification2D{
    obj: obj,
  }
}

func (me *SkeletonModification2D) BaseClass() string {
  return "SkeletonModification2D"
}

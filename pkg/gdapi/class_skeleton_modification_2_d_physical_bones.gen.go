// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonModification2DPhysicalBones struct {
  obj gdc.ObjectPtr
}

func createSkeletonModification2DPhysicalBones(obj gdc.ObjectPtr) *SkeletonModification2DPhysicalBones {
  return &SkeletonModification2DPhysicalBones{
    obj: obj,
  }
}

func (me *SkeletonModification2DPhysicalBones) BaseClass() string {
  return "SkeletonModification2DPhysicalBones"
}

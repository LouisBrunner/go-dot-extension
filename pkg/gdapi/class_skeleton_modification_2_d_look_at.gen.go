// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonModification2DLookAt struct {
  obj gdc.ObjectPtr
}

func createSkeletonModification2DLookAt(obj gdc.ObjectPtr) *SkeletonModification2DLookAt {
  return &SkeletonModification2DLookAt{
    obj: obj,
  }
}

func (me *SkeletonModification2DLookAt) BaseClass() string {
  return "SkeletonModification2DLookAt"
}

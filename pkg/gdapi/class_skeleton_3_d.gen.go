// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Skeleton3D struct {
  obj gdc.ObjectPtr
}

func createSkeleton3D(obj gdc.ObjectPtr) *Skeleton3D {
  return &Skeleton3D{
    obj: obj,
  }
}

func (me *Skeleton3D) BaseClass() string {
  return "Skeleton3D"
}
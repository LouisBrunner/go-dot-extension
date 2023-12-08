// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Skeleton2D struct {
  obj gdc.ObjectPtr
}

func createSkeleton2D(obj gdc.ObjectPtr) *Skeleton2D {
  return &Skeleton2D{
    obj: obj,
  }
}

func (me *Skeleton2D) BaseClass() string {
  return "Skeleton2D"
}

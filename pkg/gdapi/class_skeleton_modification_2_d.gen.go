// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonModification2D struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModification2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModification2D) BaseClass() string {
  return "SkeletonModification2D"
}

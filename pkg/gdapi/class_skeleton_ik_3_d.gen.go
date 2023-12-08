// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonIK3D struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonIK3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonIK3D) BaseClass() string {
  return "SkeletonIK3D"
}

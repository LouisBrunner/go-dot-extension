// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonProfileHumanoid struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonProfileHumanoid) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonProfileHumanoid) BaseClass() string {
  return "SkeletonProfileHumanoid"
}

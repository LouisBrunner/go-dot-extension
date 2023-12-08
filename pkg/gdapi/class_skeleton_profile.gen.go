// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonProfile struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonProfile) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonProfile) BaseClass() string {
  return "SkeletonProfile"
}

type SkeletonProfileTailDirection int
const (
  SkeletonProfileTailDirectionAverageChildren SkeletonProfileTailDirection = 0
  SkeletonProfileTailDirectionSpecificChild SkeletonProfileTailDirection = 1
  SkeletonProfileTailDirectionEnd SkeletonProfileTailDirection = 2
)

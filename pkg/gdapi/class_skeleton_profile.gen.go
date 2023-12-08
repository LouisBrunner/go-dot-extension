// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonProfile struct {
  obj gdc.ObjectPtr
}

func createSkeletonProfile(obj gdc.ObjectPtr) *SkeletonProfile {
  return &SkeletonProfile{
    obj: obj,
  }
}

func (me *SkeletonProfile) BaseClass() string {
  return "SkeletonProfile"
}

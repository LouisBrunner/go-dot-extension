// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonProfileHumanoid struct {
  SkeletonProfile
}

func (me *SkeletonProfileHumanoid) BaseClass() string {
  return "SkeletonProfileHumanoid"
}



// Enums

func (me *SkeletonProfileHumanoid) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonProfileHumanoid) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonProfileHumanoid) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

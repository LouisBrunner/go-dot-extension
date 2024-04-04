// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type SkeletonProfileHumanoid struct {
  SkeletonProfile
}

func (me *SkeletonProfileHumanoid) BaseClass() string {
  return "SkeletonProfileHumanoid"
}

func NewSkeletonProfileHumanoid() *SkeletonProfileHumanoid {
  str := StringNameFromStr("SkeletonProfileHumanoid") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SkeletonProfileHumanoid{}
  obj.SetBaseObject(objPtr)
  return obj
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

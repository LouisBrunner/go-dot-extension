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

type AnimationNodeBlend3 struct {
  AnimationNodeSync
}

func (me *AnimationNodeBlend3) BaseClass() string {
  return "AnimationNodeBlend3"
}

func NewAnimationNodeBlend3() *AnimationNodeBlend3 {
  str := StringNameFromStr("AnimationNodeBlend3") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeBlend3{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AnimationNodeBlend3) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeBlend3) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlend3) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

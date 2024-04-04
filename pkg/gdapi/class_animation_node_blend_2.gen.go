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

type AnimationNodeBlend2 struct {
  AnimationNodeSync
}

func (me *AnimationNodeBlend2) BaseClass() string {
  return "AnimationNodeBlend2"
}

func NewAnimationNodeBlend2() *AnimationNodeBlend2 {
  str := StringNameFromStr("AnimationNodeBlend2") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeBlend2{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AnimationNodeBlend2) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeBlend2) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlend2) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

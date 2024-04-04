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

type AnimationNodeTimeSeek struct {
  AnimationNode
}

func (me *AnimationNodeTimeSeek) BaseClass() string {
  return "AnimationNodeTimeSeek"
}

func NewAnimationNodeTimeSeek() *AnimationNodeTimeSeek {
  str := StringNameFromStr("AnimationNodeTimeSeek") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeTimeSeek{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AnimationNodeTimeSeek) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeTimeSeek) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeTimeSeek) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

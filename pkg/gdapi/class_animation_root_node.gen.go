// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationRootNode struct {
  AnimationNode
}

func (me *AnimationRootNode) BaseClass() string {
  return "AnimationRootNode"
}

func NewAnimationRootNode() *AnimationRootNode {
  str := StringNameFromStr("AnimationRootNode") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationRootNode{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AnimationRootNode) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationRootNode) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationRootNode) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeSub2 struct {
  AnimationNodeSync
}

func (me *AnimationNodeSub2) BaseClass() string {
  return "AnimationNodeSub2"
}



// Enums

func (me *AnimationNodeSub2) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeSub2) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeSub2) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

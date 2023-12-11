// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeAdd2 struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeAdd2) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeAdd2) BaseClass() string {
  return "AnimationNodeAdd2"
}



// Enums

func (me *AnimationNodeAdd2) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeAdd2) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeAdd2) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

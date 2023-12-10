// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeOutput struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeOutput) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeOutput) BaseClass() string {
  return "AnimationNodeOutput"
}



// Enums

func (me *AnimationNodeOutput) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeOutput) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeOutput) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties
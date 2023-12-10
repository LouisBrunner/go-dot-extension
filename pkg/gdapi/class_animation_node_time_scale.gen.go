// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeTimeScale struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeTimeScale) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeTimeScale) BaseClass() string {
  return "AnimationNodeTimeScale"
}



// Enums

func (me *AnimationNodeTimeScale) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeTimeScale) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeTimeScale) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties
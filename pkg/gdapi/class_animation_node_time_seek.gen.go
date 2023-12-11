// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeTimeSeek struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeTimeSeek) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeTimeSeek) BaseClass() string {
  return "AnimationNodeTimeSeek"
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

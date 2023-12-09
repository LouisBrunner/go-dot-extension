// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

// TODO: properties (class)

// TODO: signals (class)

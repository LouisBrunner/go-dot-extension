// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeBlend3 struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeBlend3) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeBlend3) BaseClass() string {
  return "AnimationNodeBlend3"
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

// TODO: properties (class)

// TODO: signals (class)

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeBlend2 struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeBlend2) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeBlend2) BaseClass() string {
  return "AnimationNodeBlend2"
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

// TODO: properties (class)

// TODO: signals (class)

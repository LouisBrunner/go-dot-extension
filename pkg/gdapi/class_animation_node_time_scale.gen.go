// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

// TODO: properties (class)

// TODO: signals (class)

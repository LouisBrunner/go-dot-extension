// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationRootNode struct {
  obj gdc.ObjectPtr
}

func (me *AnimationRootNode) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationRootNode) BaseClass() string {
  return "AnimationRootNode"
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

// TODO: properties (class)

// TODO: signals (class)

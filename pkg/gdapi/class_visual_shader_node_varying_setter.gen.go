// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVaryingSetter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVaryingSetter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVaryingSetter) BaseClass() string {
  return "VisualShaderNodeVaryingSetter"
}



// Enums

func (me *VisualShaderNodeVaryingSetter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVaryingSetter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVaryingSetter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)

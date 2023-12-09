// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeRemap struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeRemap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeRemap) BaseClass() string {
  return "VisualShaderNodeRemap"
}



// Enums

func (me *VisualShaderNodeRemap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeRemap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeRemap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)

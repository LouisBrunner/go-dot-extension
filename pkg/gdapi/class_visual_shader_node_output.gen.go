// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeOutput struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeOutput) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeOutput) BaseClass() string {
  return "VisualShaderNodeOutput"
}



// Enums

func (me *VisualShaderNodeOutput) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeOutput) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeOutput) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)

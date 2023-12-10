// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeGlobalExpression struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeGlobalExpression) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeGlobalExpression) BaseClass() string {
  return "VisualShaderNodeGlobalExpression"
}



// Enums

func (me *VisualShaderNodeGlobalExpression) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeGlobalExpression) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeGlobalExpression) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties
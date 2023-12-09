// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeLinearSceneDepth struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeLinearSceneDepth) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeLinearSceneDepth) BaseClass() string {
  return "VisualShaderNodeLinearSceneDepth"
}



// Enums

func (me *VisualShaderNodeLinearSceneDepth) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeLinearSceneDepth) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeLinearSceneDepth) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)

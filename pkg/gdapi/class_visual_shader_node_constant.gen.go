// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeConstant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeConstant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeConstant) BaseClass() string {
  return "VisualShaderNodeConstant"
}



// Enums

func (me *VisualShaderNodeConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)

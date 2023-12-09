// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformParameter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformParameter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformParameter) BaseClass() string {
  return "VisualShaderNodeTransformParameter"
}



// Enums

func (me *VisualShaderNodeTransformParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeTransformParameter) SetDefaultValueEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTransformParameter) IsDefaultValueEnabled()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTransformParameter) SetDefaultValue(value Transform3D, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTransformParameter) GetDefaultValue()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVec3Parameter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVec3Parameter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVec3Parameter) BaseClass() string {
  return "VisualShaderNodeVec3Parameter"
}



// Enums

func (me *VisualShaderNodeVec3Parameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec3Parameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeVec3Parameter) SetDefaultValueEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVec3Parameter) IsDefaultValueEnabled()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVec3Parameter) SetDefaultValue(value Vector3, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVec3Parameter) GetDefaultValue()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

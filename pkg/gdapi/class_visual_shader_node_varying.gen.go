// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVarying struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVarying) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVarying) BaseClass() string {
  return "VisualShaderNodeVarying"
}



// Enums

func (me *VisualShaderNodeVarying) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVarying) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVarying) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeVarying) SetVaryingName(name String, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVarying) GetVaryingName()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVarying) SetVaryingType(type_ VisualShaderVaryingType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVarying) GetVaryingType()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

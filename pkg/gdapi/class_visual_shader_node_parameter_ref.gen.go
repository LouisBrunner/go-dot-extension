// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParameterRef struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParameterRef) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParameterRef) BaseClass() string {
  return "VisualShaderNodeParameterRef"
}



// Enums

func (me *VisualShaderNodeParameterRef) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParameterRef) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParameterRef) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeParameterRef) SetParameterName(name String, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeParameterRef) GetParameterName()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

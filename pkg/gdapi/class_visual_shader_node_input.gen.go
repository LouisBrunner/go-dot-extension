// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeInput struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeInput) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeInput) BaseClass() string {
  return "VisualShaderNodeInput"
}



// Enums

func (me *VisualShaderNodeInput) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeInput) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeInput) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeInput) SetInputName(name String, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeInput) GetInputName()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeInput) GetInputRealName()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

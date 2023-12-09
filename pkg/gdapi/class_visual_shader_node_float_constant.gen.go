// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFloatConstant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeFloatConstant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeFloatConstant) BaseClass() string {
  return "VisualShaderNodeFloatConstant"
}



// Enums

func (me *VisualShaderNodeFloatConstant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFloatConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFloatConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeFloatConstant) SetConstant(constant float32, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatConstant) GetConstant()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

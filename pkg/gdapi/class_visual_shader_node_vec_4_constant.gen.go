// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVec4Constant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVec4Constant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVec4Constant) BaseClass() string {
  return "VisualShaderNodeVec4Constant"
}



// Enums

func (me *VisualShaderNodeVec4Constant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVec4Constant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec4Constant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeVec4Constant) SetConstant(constant Quaternion, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVec4Constant) GetConstant()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

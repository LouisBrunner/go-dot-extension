// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVec2Constant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVec2Constant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVec2Constant) BaseClass() string {
  return "VisualShaderNodeVec2Constant"
}



// Enums

func (me *VisualShaderNodeVec2Constant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVec2Constant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec2Constant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeVec2Constant) SetConstant(constant Vector2, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVec2Constant) GetConstant()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

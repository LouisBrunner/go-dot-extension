// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVec3Constant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVec3Constant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVec3Constant) BaseClass() string {
  return "VisualShaderNodeVec3Constant"
}



// Enums

func (me *VisualShaderNodeVec3Constant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec3Constant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeVec3Constant) SetConstant(constant Vector3, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVec3Constant) GetConstant()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

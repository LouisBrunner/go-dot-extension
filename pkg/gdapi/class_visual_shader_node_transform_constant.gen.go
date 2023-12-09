// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformConstant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformConstant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformConstant) BaseClass() string {
  return "VisualShaderNodeTransformConstant"
}



// Enums

func (me *VisualShaderNodeTransformConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeTransformConstant) SetConstant(constant Transform3D, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTransformConstant) GetConstant()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)

// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeCurveTexture struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeCurveTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeCurveTexture) BaseClass() string {
  return "VisualShaderNodeCurveTexture"
}



// Enums

func (me *VisualShaderNodeCurveTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCurveTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeCurveTexture) SetTexture(texture CurveTexture, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCurveTexture) GetTexture()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
